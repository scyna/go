package gateway

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/scyna/go/scyna"
	"google.golang.org/protobuf/proto"
)

type Gateway struct {
	Queries      QueryPool
	Applications map[string]Application
	Contexts     scyna.HttpContextPool
}

func NewGateway() *Gateway {
	ret := &Gateway{
		Queries:  NewQueryPool(),
		Contexts: scyna.NewContextPool(),
	}
	ret.initApplications()
	return ret
}

func (gateway *Gateway) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var app Application
	callID := scyna.ID.Next()

	auth := false
	ok, appID, json, url := parseUrl(req.URL.String())

	if !ok {
		log.Print("Path not ok")
		http.Error(rw, "NotFound", http.StatusNotFound)
		return
	}

	log.Print(url)

	query := gateway.Queries.GetQuery()
	defer gateway.Queries.Put(query)

	ctx := gateway.Contexts.GetContext()
	defer gateway.Contexts.PutContext(ctx)

	trace := scyna.Trace{
		ID:       callID,
		ParentID: 0,
		Time:     time.Now(),
		Path:     url,
		Type:     scyna.TRACE_SERVICE,
		Source:   app.Code,
	}
	defer trace.Save()

	/*headers*/
	rw.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Access-Control-Allow-Methods", "POST")
	if json {
		rw.Header().Set("Content-Type", "application/json")
	} else {
		rw.Header().Set("Content-Type", "application/octet-stream")
	}

	ctx.Request.JSON = json
	ctx.Request.TraceID = callID

	if a, ok := gateway.Applications[appID]; !ok {
		http.Error(rw, "Forbidden", http.StatusForbidden)
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusForbidden
		return
	} else {
		app = a
	}

	if url == "/auth" {
		auth = true
		url = app.AuthURL
		log.Print(url)
	} else {
		if cookie, err := req.Cookie("session"); err == nil {
			token := cookie.Value
			ctx.Request.Data = token
			if exp := checkService(token, appID, url); exp == nil {
				http.Error(rw, "Unauthorized", http.StatusUnauthorized)
				trace.SessionID = scyna.Session.ID()
				trace.Status = http.StatusUnauthorized
				return
			} else {
				//log.Print(exp)
				now := time.Now()
				if exp.Before(now) {
					http.Error(rw, "Unauthorized", http.StatusUnauthorized)
					scyna.LOG.Info("Session expired")
					trace.SessionID = scyna.Session.ID()
					trace.Status = http.StatusUnauthorized
					return
				} else {
					if exp.After(now.Add(time.Minute * 10)) {
						/*auto extend expire*/
						if updateSession(token, now.Add(time.Hour*8)) {
							cookie.Expires = now
							http.SetCookie(rw, cookie)
						}
					}
				}
			}
		} else {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			scyna.LOG.Info("Can not get cookie")
			trace.SessionID = scyna.Session.ID()
			trace.Status = http.StatusUnauthorized
			return
		}
	}

	/*build request*/
	err := ctx.Request.Build(req)
	if err != nil {
		http.Error(rw, "Cannot process request", http.StatusInternalServerError)
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusInternalServerError
		return
	}

	/*serialize the request */
	reqBytes, err := proto.Marshal(&ctx.Request)
	if err != nil {
		http.Error(rw, "Cannot process request", http.StatusInternalServerError)
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusInternalServerError
		return
	}

	/*post request to message queue*/
	var _url string
	if url != app.AuthURL {
		_url = "/" + appID + url
	} else {
		_url = url
	}
	msg, respErr := scyna.Connection.Request(scyna.PublishURL(_url), reqBytes, 10*time.Second)
	if respErr != nil {
		http.Error(rw, "No response", http.StatusInternalServerError)
		scyna.LOG.Error("ServeHTTP: Nats: " + respErr.Error())
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusInternalServerError
		return
	}

	/*response*/
	if err := proto.Unmarshal(msg.Data, &ctx.Response); err != nil {
		log.Println()
		http.Error(rw, "Cannot deserialize response", http.StatusInternalServerError)
		scyna.LOG.Error("nats-proxy:" + err.Error())
		trace.SessionID = scyna.Session.ID()
		trace.Status = http.StatusInternalServerError
		return
	}

	if auth {
		if ctx.Response.Code == 200 {
			cookie := &http.Cookie{
				Name:     "session",
				Value:    ctx.Response.Token,
				Path:     "/",
				Expires:  time.Unix(0, int64(ctx.Response.Expired*1000)),
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			}
			http.SetCookie(rw, cookie)
			log.Print("Set cookie:", ctx.Response.Token)
		} else {
			c := &http.Cookie{
				Name:     "session",
				Value:    "",
				Path:     "/",
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			}
			http.SetCookie(rw, c) /*clear cookie*/
			/*TODO: make Authentication inactive here or delete from database*/
		}
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
