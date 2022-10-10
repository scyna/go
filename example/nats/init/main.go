package main

import (
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "nats"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		Url:     "https://localhost:8081",
		Context: MODULE_CODE,
		Secret:  "123456",
	})
	defer scyna.Release()

	utils.DeleteStream("nats_stream")
	utils.CreateStreamForModule("nats_stream")
	utils.CreateSyncConsumer("nats_stream", "channel1", "receiver1")
	utils.CreateSyncConsumer("nats_stream", "channel1", "receiver2")
	scyna.Start()
}
