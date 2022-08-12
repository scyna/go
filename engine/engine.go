package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/scyna/go/engine/gateway"
	"github.com/scyna/go/engine/manager/authentication"
	"github.com/scyna/go/engine/manager/generator"
	"github.com/scyna/go/engine/manager/logging"
	"github.com/scyna/go/engine/manager/scheduler"
	"github.com/scyna/go/engine/manager/session"
	"github.com/scyna/go/engine/manager/setting"
	"github.com/scyna/go/engine/manager/trace"
	"github.com/scyna/go/engine/proxy"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "scyna.engine"

func main() {
	managerPort := flag.String("manager_port", "8081", "Manager Port")
	proxyPort := flag.String("proxy_port", "8080", "Proxy Port")
	gatewayPort := flag.String("gateway_port", "8443", "GateWay Port")

	natsUrl := flag.String("nats_url", "127.0.0.1", "Nats URL")
	natsUsername := flag.String("nats_username", "", "Nats Username")
	natsPassword := flag.String("nats_password", "", "Nats Password")
	dbHost := flag.String("db_host", "127.0.0.1", "DB Host")
	dbUsername := flag.String("db_username", "", "DB Username")
	dbPassword := flag.String("db_password", "", "DB Password")
	dbLocation := flag.String("db_location", "", "DB Location")
	secret := flag.String("secret", "123456", "scyna Manager Secret")

	certificateFile := flag.String("certificateFile", "", "Certificate Key")
	certificateKey := flag.String("certificateKey", "", "Certificate File")

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
	scyna.RegisterSignalLite(scyna.TRACE_CREATED_CHANNEL, trace.TraceCreated)
	scyna.RegisterSignalLite(scyna.TAG_CREATED_CHANNEL, trace.TagCreated)
	scyna.RegisterSignalLite(scyna.SERVICE_DONE_CHANNEL, trace.ServiceDone)

	/*setting*/
	scyna.RegisterService(scyna.SETTING_READ_URL, setting.Read)
	scyna.RegisterService(scyna.SETTING_WRITE_URL, setting.Write)
	scyna.RegisterService(scyna.SETTING_REMOVE_URL, setting.Remove)

	/*authentication*/
	scyna.RegisterService(scyna.AUTH_CREATE_URL, authentication.Create)
	scyna.RegisterService(scyna.AUTH_GET_URL, authentication.Get)
	scyna.RegisterService(scyna.AUTH_LOGOUT_URL, authentication.Logout)

	/* task */
	scyna.RegisterService(scyna.START_TASK_URL, scheduler.StartTask)
	scyna.RegisterService(scyna.STOP_TASK_URL, scheduler.StopTask)

	/* Update config */
	setting.UpdateDefaultConfig(&config)

	const DEFAULT_CERT_FILE = ".cert/localhost.crt"
	const DEFAULT_CERT_KEY = ".cert/localhost.key"

	if *certificateFile == "" || *certificateKey == "" {
		*certificateFile = DEFAULT_CERT_FILE
		*certificateKey = DEFAULT_CERT_KEY
	}

	go func() {
		gateway_ := gateway.NewGateway()
		log.Println("Scyna Gateway Start with port " + *gatewayPort)
		if err := http.ListenAndServeTLS(":"+*gatewayPort, *certificateFile, *certificateKey, gateway_); err != nil {
			log.Println("Gateway: " + err.Error())
		}
	}()

	go func() {
		proxy_ := proxy.NewProxy()
		log.Println("Scyna Proxy Start with port " + *proxyPort)
		if err := http.ListenAndServeTLS(":"+*proxyPort, *certificateFile, *certificateKey, proxy_); err != nil {
			log.Println("Proxy: " + err.Error())
		}
	}()

	/* Start worker */
	scheduler.Start(time.Second * 10)
	/*session*/
	scyna.RegisterSignalLite(scyna.SESSION_END_CHANNEL, session.End)
	scyna.RegisterSignalLite(scyna.SESSION_UPDATE_CHANNEL, session.Update)
	http.HandleFunc(scyna.SESSION_CREATE_URL, session.Create)
	log.Println("Scyna Manager Start with port " + *managerPort)
	log.Fatal(http.ListenAndServe(":"+*managerPort, nil))
	// log.Fatal(http.ListenAndServeTLS(":"+*managerPort, *certificateFile, *certificateKey, nil))
}
