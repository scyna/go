package proxy

import (
	"bytes"
	"github.com/scyna/go/scyna"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"strings"
	"time"
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

	if !ok || clientSecret != client.Secret {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		log.Print("Wrong client id or secret: ", clientID)
		return
	}

	// if client.State != uint32(scyna.ClientState_ACTIVE) {
	// 	http.Error(rw, "Unauthorized", http.StatusUnauthorized)
	// 	log.Printf("Client is inactive: %s\n", clientID)
	// 	proxy.SaveErrorCall(clientID, 401, callID, day, start, req.URL.Path)
	// 	return
	// }

	query := proxy.Queries.GetQuery()
	defer proxy.Queries.Put(query)
	ctx := proxy.Contexts.GetContext()
	defer proxy.Contexts.PutContext(ctx)

	if err := query.Authenticate.Bind(clientID, url).Get(&url); err != nil {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		log.Printf("Wrong url: %s, error = %s\n", url, err.Error())
		return
	}

	if contentType == "application/json" {
		ctx.Request.JSON = true
	} else if contentType == "application/protobuf" {
		ctx.Request.JSON = false
	} else {
		http.Error(rw, "Content-Type must be JSON or PROTOBUF ", http.StatusNotAcceptable)
		return
	}

	context := scyna.Context{
		ID:        callID,
		ParentID:  0,
		Time:      time.Now(),
		Path:      url,
		Type:      scyna.TRACE_SERVICE,
		Source:    clientID,
		SessionID: scyna.Session.ID(),
	}
	defer proxy.saveContext(&context)

	/*build request*/
	err := ctx.Request.Build(req)
	if err != nil {
		http.Error(rw, "Cannot process request", http.StatusInternalServerError)
		return
	}

	ctx.Request.TraceID = callID
	ctx.Request.Data = client.Type

	/*serialize the request */
	reqBytes, err := proto.Marshal(&ctx.Request)
	if err != nil {
		http.Error(rw, "Cannot process request", http.StatusInternalServerError)
		context.Status = http.StatusInternalServerError
		return
	}

	/*post request to message queue*/
	msg, respErr := scyna.Connection.Request(scyna.PublishURL(url), reqBytes, 10*time.Second)
	if respErr != nil {
		http.Error(rw, "No response", http.StatusInternalServerError)
		context.Status = http.StatusInternalServerError
		log.Println("ServeHTTP: Nats: " + respErr.Error())
		return
	}

	/*response*/
	if err := proto.Unmarshal(msg.Data, &ctx.Response); err != nil {
		log.Println("nats-proxy:" + err.Error())
		http.Error(rw, "Cannot deserialize response", http.StatusInternalServerError)
		context.Status = http.StatusInternalServerError
		return
	}

	rw.WriteHeader(int(ctx.Response.Code))
	_, err = bytes.NewBuffer(ctx.Response.Body).WriteTo(rw)
	if err != nil {
		log.Println("Proxy write data error: " + err.Error())
	}

	if f, ok := rw.(http.Flusher); ok {
		f.Flush()
	}

	context.Status = ctx.Response.Code
}
