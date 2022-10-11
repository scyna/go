package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/scyna/go/engine/manager/generator"
	"github.com/scyna/go/engine/manager/logging"
	"github.com/scyna/go/engine/manager/session"
	"github.com/scyna/go/scyna"
)

const CONTEXT_CODE = "scyna.engine"

func main() {
	managerPort := flag.String("manager_port", "8081", "Manager Port")
	natsUrl := flag.String("nats_url", "127.0.0.1", "Nats URL")
	natsUsername := flag.String("nats_username", "", "Nats Username")
	natsPassword := flag.String("nats_password", "", "Nats Password")
	dbHost := flag.String("db_host", "127.0.0.1", "DB Host")
	dbUsername := flag.String("db_username", "", "DB Username")
	dbPassword := flag.String("db_password", "", "DB Password")
	dbLocation := flag.String("db_location", "", "DB Location")
	secret := flag.String("secret", "123456", "scyna Manager Secret")

	flag.Parse()
	config := scyna.Configuration{
		NatsUrl:      *natsUrl,
		NatsUsername: *natsUsername,
		NatsPassword: *natsPassword,
		DBHost:       *dbHost,
		DBUsername:   *dbUsername,
		DBPassword:   *dbPassword,
		DBLocation:   *dbLocation,
	}

	/* Init module */
	scyna.DirectInit(&config)
	defer scyna.Release()
	generator.Init()
	session.Init(CONTEXT_CODE, *secret)
	scyna.UseDirectLog(5)

	/* generator */
	scyna.RegisterService(scyna.GEN_GET_ID_URL, generator.GetID)
	scyna.RegisterService(scyna.GEN_GET_SN_URL, generator.GetSN)

	/*logging*/
	scyna.RegisterSignal(scyna.LOG_CREATED_CHANNEL, logging.Write)

	/*session*/
	scyna.RegisterSignal(scyna.SESSION_END_CHANNEL, session.End)
	scyna.RegisterSignal(scyna.SESSION_UPDATE_CHANNEL, session.Update)
	http.HandleFunc(scyna.SESSION_CREATE_URL, session.Create)
	log.Println("Scyna Manager Start with port " + *managerPort)
	log.Fatal(http.ListenAndServe(":"+*managerPort, nil))
}
