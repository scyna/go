package proxy

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/scyna/go/scyna"
	"google.golang.org/protobuf/proto"
)

type Proxy struct {
	Queries  QueryPool
	Clients  map[string]Client
	Contexts scyna.ContextPool
}

func NewProxy() *Proxy {
	ret := &Proxy{
		Queries:  NewQueryPool(),
		Contexts: scyna.NewContextPool(),
	}
	ret.initClients()
	return ret
}

func (proxy *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	callID := scyna.ID.Next()

	/*authenticate*/
	url := req.URL.String()
	clientID := req.Header.Get("Client-Id")
	clientSecret := req.Header.Get("Client-Secret")
	client, ok := proxy.Clients[clientID]
	contentType := req.Header.Get("Content-Type")

	//https://descynaper.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
	for _, data := range strings.Split(contentType, ";") {
		value := strings.TrimSpace(strings.Trim(data, ";"))
		if strings.HasPrefix(value, "application/") {
			contentType = value
			continue
		}
	}

	/*CORS*/
	rw.Header().Set("Content-Type", contentType)
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	trace := scyna.Trace{
		ID:       callID,
		ParentID: 0,
		Time:     time.Now(),
		Path:     url,
		Type:     scyna.TRACE_SERVICE,
		Source:   clientID,
	}
	defer trace.Save()

	query := proxy.Queries.GetQuery()
	defer proxy.Queries.Put(query)
	ctx := proxy.Contexts.GetContext()
	defer proxy.Contexts.PutContext(ctx)

	if !ok || clientSecret != client.Secret {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		scyna.LOG.Info("Wrong client id or secret: " + clientID)
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusUnauthorized
		return
	}

	// if client.State != uint32(scyna.ClientState_ACTIVE) {
	// 	http.Error(rw, "Unauthorized", http.StatusUnauthorized)
	// 	log.Printf("Client is inactive: %s\n", clientID)
	// 	proxy.SaveErrorCall(clientID, 401, callID, day, start, req.URL.Path)
	// 	return
	// }

	if err := query.Authenticate.Bind(clientID, url).Get(&url); err != nil {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		scyna.LOG.Info(fmt.Sprintf("Wrong url: %s, error = %s\n", url, err.Error()))
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusUnauthorized
		return
	}

	if contentType == "application/json" {
		ctx.Request.JSON = true
	} else if contentType == "application/protobuf" {
		ctx.Request.JSON = false
	} else {
		http.Error(rw, "Content-Type must be JSON or PROTOBUF ", http.StatusNotAcceptable)
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusNotAcceptable
		return
	}

	/*build request*/
	err := ctx.Request.Build(req)
	if err != nil {
		http.Error(rw, "Cannot process request", http.StatusInternalServerError)
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusInternalServerError
		return
	}

	ctx.Request.TraceID = callID
	ctx.Request.Data = client.Type

	/*serialize the request */
	reqBytes, err := proto.Marshal(&ctx.Request)
	if err != nil {
		http.Error(rw, "Cannot process request", http.StatusInternalServerError)
		trace.Status = http.StatusInternalServerError
		trace.SessionID = scyna.Session.ID()
		return
	}

	/*post request to message queue*/
	msg, respErr := scyna.Connection.Request(scyna.PublishURL(url), reqBytes, 10*time.Second)
	if respErr != nil {
		http.Error(rw, "No response", http.StatusInternalServerError)
		trace.Status = http.StatusInternalServerError
		trace.SessionID = scyna.Session.ID()
		scyna.LOG.Error("ServeHTTP: Nats: " + respErr.Error())
		return
	}

	/*response*/
	if err := proto.Unmarshal(msg.Data, &ctx.Response); err != nil {
		http.Error(rw, "Cannot deserialize response", http.StatusInternalServerError)
		scyna.LOG.Error("nats-proxy:" + err.Error())
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusInternalServerError
		return
	}

	rw.WriteHeader(int(ctx.Response.Code))
	trace.SessionID = ctx.Response.SessionID
	trace.Status = ctx.Response.Code
	_, err = bytes.NewBuffer(ctx.Response.Body).WriteTo(rw)
	if err != nil {
		scyna.LOG.Error("Proxy write data error: " + err.Error())
		trace.SessionID = scyna.Session.ID()
		trace.Status = 0

	}

	if f, ok := rw.(http.Flusher); ok {
		f.Flush()
	}
}
