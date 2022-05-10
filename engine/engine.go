package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/scyna/engine/gateway"
	"github.com/scyna/engine/manager/authentication"
	"github.com/scyna/engine/manager/generator"
	"github.com/scyna/engine/manager/logging"
	"github.com/scyna/engine/manager/session"
	"github.com/scyna/engine/manager/setting"
	"github.com/scyna/engine/manager/trace"
	scyna "github.com/scyna/go/scyna"
)

const MODULE_CODE = "scyna.engine"

func main() {
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
	scyna.DirectInit(MODULE_CODE, &config)
	defer scyna.Release()
	generator.Init()
	session.Init(MODULE_CODE, *secret)
	scyna.UseDirectLog(5)

	/* generator */
	scyna.RegisterCommand(scyna.GEN_GET_ID_URL, generator.GetID)
	scyna.RegisterService(scyna.GEN_GET_SN_URL, generator.GetSN)

	/*logging*/
	scyna.RegisterSignalLite(scyna.LOG_CREATED_CHANNEL, logging.Write)

	/*trace*/
	scyna.RegisterSignalLite(scyna.TRACE_CREATED_CHANNEL, trace.WriteTrace)
	scyna.RegisterSignalLite(scyna.TAG_CREATED_CHANNEL, trace.WriteTag)

	/*setting*/
	scyna.RegisterService(scyna.SETTING_READ_URL, setting.Read)
	scyna.RegisterService(scyna.SETTING_WRITE_URL, setting.Write)
	scyna.RegisterService(scyna.SETTING_REMOVE_URL, setting.Remove)

	/*authentication*/
	scyna.RegisterService(scyna.AUTH_CREATE_URL, authentication.Create)
	scyna.RegisterService(scyna.AUTH_GET_URL, authentication.Get)
	scyna.RegisterService(scyna.AUTH_LOGOUT_URL, authentication.Logout)

	/* Update config */
	setting.UpdateDefautConfig(&config)

	go func() {
		gateway_ := gateway.NewGateway()
		log.Println("Scyna Gateway Started")
		if err := http.ListenAndServe(":8443", gateway_); err != nil {
			log.Println("Gateway:" + err.Error())
		}
	}()

	//go func() {
	//	proxy_ := proxy.NewProxy()
	//	log.Println("Scyna Proxy Started")
	//	if err := http.ListenAndServe(":8080", proxy_); err != nil {
	//		log.Println("Proxy:" + err.Error())
	//	}
	//}()

	/*session*/
	scyna.RegisterSignalLite(scyna.SESSION_END_CHANNEL, session.End)
	scyna.RegisterSignalLite(scyna.SESSION_UPDATE_CHANNEL, session.Update)
	http.HandleFunc(scyna.SESSION_CREATE_URL, session.Create)
	log.Println("Scyna Manager Started")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
