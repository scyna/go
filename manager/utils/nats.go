package utils

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/scyna/go/scyna"
)

func DeleteStream(name string) error {
	err := scyna.JetStream.DeleteStream(name)
	if err != nil {
		return err
	}
	return nil
}

func CreateStreamForModule(module string) error {
	if _, err := scyna.JetStream.AddStream(&nats.StreamConfig{
		Name:     module,
		Subjects: []string{module + ".>"},
		Storage:  nats.FileStorage,
		MaxAge:   time.Hour * 24 * 7, //keep for a week
	}); err != nil {
		log.Print("Create stream for module " + module + ": Error: " + err.Error())
		return err
	}
	return nil
}

func CreateConsumer(source string, target string) error {

	/*check if consumer exists*/
	if _, err := scyna.JetStream.ConsumerInfo(source, target); err != nil {
		return nil
	}

	/*create pull consumer*/
	if _, err := scyna.JetStream.AddConsumer(source, &nats.ConsumerConfig{
		Durable: target,
	}); err != nil {
		log.Print("Add Consumer "+target+": Error: ", err)
		return err
	}
	return nil
}

func CreateSyncConsumer(module string, channel string) error {
	consumerName := "sync_" + channel

	/*check if consumer exists*/
	if _, err := scyna.JetStream.ConsumerInfo(module, consumerName); err != nil {
		return nil
	}

	/*create push consumer*/
	if _, err := scyna.JetStream.AddConsumer(module, &nats.ConsumerConfig{
		Durable:        consumerName,
		DeliverSubject: ".sync." + channel,
		FilterSubject:  module + ".sync." + channel,
	}); err != nil {
		log.Print("Add Consumer "+consumerName+": Error: ", err)
		return err
	}
	return nil
}
