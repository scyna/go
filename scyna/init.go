package scyna

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/protobuf/proto"
)

type RemoteConfig struct {
	Url     string
	Context string
	Secret  string
}

func RemoteInit(config RemoteConfig) {

	request := CreateSessionRequest{
		Context: config.Context,
		Secret:  config.Secret,
	}

	data, err := proto.Marshal(&request)
	if err != nil {
		log.Fatal("Bad authentication request")
	}

	req, err := http.NewRequest("POST", config.Url+SESSION_CREATE_URL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in create http request:", err)
	}

	res, err := HttpClient().Do(req)
	if err != nil {
		log.Fatal("Error in send http request:", err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Error in autheticate")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Can not read response body:", err)
	}

	var response CreateSessionResponse
	if err := proto.Unmarshal(resBody, &response); err != nil {
		log.Fatal("Authenticate error")
	}

	Session = NewSession(response.SessionID, config.Context)
	initNats(response.Config)
	initScylla(response.Config)
}

func DirectInit(c *Configuration) {
	initNats(c)
	initScylla(c)
}

func initNats(c *Configuration) {
	var nats_ []string
	var err error

	for _, n := range strings.Split(c.NatsUrl, ",") {
		fmt.Printf("Nats configuration: nats://%s:4222\n", n)
		nats_ = append(nats_, fmt.Sprintf("nats://%s:4222", n))
	}

	if c.NatsUsername != "" && c.NatsPassword != "" {
		Connection, err = nats.Connect(strings.Join(nats_, ","), nats.UserInfo(c.NatsUsername, c.NatsPassword))
	} else {
		Connection, err = nats.Connect(strings.Join(nats_, ","))
	}

	if err != nil {
		log.Fatal("Can not connect to NATS:", nats_)
	}

	JetStream, err = Connection.JetStream()
	if err != nil {
		log.Fatal("Init: " + err.Error())
	}
}

func initScylla(c *Configuration) {
	hosts := strings.Split(c.DBHost, ",")
	cluster := gocql.NewCluster(hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: c.DBUsername, Password: c.DBPassword}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(c.DBLocation)
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	cluster.Consistency = gocql.Quorum

	log.Printf("Connecting to: %s\n", hosts)

	var err error
	DB, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalf("Can not create session: Host = %s, Error = %s ", hosts, err.Error())
	}
}
