package utils

import (
	"errors"
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
		Name:         module,
		Subjects:     []string{module + ".>"},
		Storage:      nats.FileStorage,
		MaxAge:       time.Hour * 24 * 7, //keep for a week
		Replicas:     3,
		Retention:    nats.LimitsPolicy,
		MaxMsgs:      -1,
		MaxConsumers: -1,
		MaxBytes:     -1,
		MaxMsgSize:   -1,
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
		Durable:       target,
		FilterSubject: source + ".*",
	}); err != nil {
		log.Print("Add Consumer "+target+": Error: ", err)
		return err
	}
	return nil
}

func CreateSyncConsumer(module string, channel string, receiver string) error {

	consumerName := "sync_" + channel + "_" + receiver

	/*check if consumer exists*/
	if _, err := scyna.JetStream.ConsumerInfo(module, consumerName); err == nil {
		return errors.New("consumer existed")
	}

	/*create pull consumer*/
	if _, err := scyna.JetStream.AddConsumer(module, &nats.ConsumerConfig{
		Durable:       consumerName,
		FilterSubject: module + ".sync." + channel,
		AckPolicy:     nats.AckExplicitPolicy,
	}); err != nil {
		log.Print("Add Consumer "+consumerName+": Error: ", err)
		return err
	}
	return nil
}

func CreateTaskConsumer(module string, task string) error {
	consumerName := "task_" + task

	/*check if consumer exists*/
	if _, err := scyna.JetStream.ConsumerInfo(module, consumerName); err == nil {
		return errors.New("consumer existed")
	}

	/*create pull consumer*/
	if _, err := scyna.JetStream.AddConsumer(module, &nats.ConsumerConfig{
		Durable:       consumerName,
		FilterSubject: module + ".task." + task,
		AckPolicy:     nats.AckExplicitPolicy,
	}); err != nil {
		log.Print("Add Consumer "+consumerName+": Error: ", err)
		return err
	}
	return nil
}
