package session

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/scyna/go/engine/manager/manager"
	"github.com/scyna/go/scyna"
	"google.golang.org/protobuf/proto"
)

func Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Receive CreateSessionRequest")
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request scyna.CreateSessionRequest
	if err := proto.Unmarshal(buf, &request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if request.Module == manager.MODULE_CODE {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if sid, err := newSession(request.Module, request.Secret); err == scyna.OK {
		var response scyna.CreateSessionResponse
		response.SessionID = sid

		response.Config = manager.DefaultConfig

		if data, err := proto.Marshal(&response); err == nil {
			w.WriteHeader(200)
			_, err = bytes.NewBuffer(data).WriteTo(w)
			if err != nil {
				log.Println("Proxy write data error: " + err.Error())
			}
		} else {
			http.Error(w, "Server Error", 400)
		}
		return
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
