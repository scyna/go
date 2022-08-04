package scyna

import (
	"crypto/tls"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2"
)

const REQUEST_TIMEOUT = 10

var Connection *nats.Conn
var JetStream nats.JetStreamContext
var Session *session
var DB gocqlx.Session
var ID generator
var Settings settings

var httpClient *http.Client
var module string
var LOG *Logger

func Release() {
	releaseLog()
	Session.release()
	Connection.Close()
	DB.Close()
}

func Start() {
	startEventStream()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func HttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}
	return httpClient
}
