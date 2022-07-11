package utils

import (
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/scyna/go/scyna"
	"log"
	"time"
)

func DeleteStream(name string) error {
	err := scyna.JetStream.DeleteStream(name)
	if err != nil {
		return err
	}
	return nil
}

func CreateStream(name string) error {
	if _, err := scyna.JetStream.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".*"},
		Storage:  nats.FileStorage,
		MaxAge:   time.Hour * 24 * 7, //keep for a week
	}); err != nil {
		return errors.New("CreateStream " + name + ": Error: " + err.Error())
	}
	return nil
}

func CreateConsumer(stream string, name string, channel string) error {
	if _, err := scyna.JetStream.AddConsumer(stream, &nats.ConsumerConfig{
		Durable:      channel,
		DeliverGroup: name,
		//DeliverSubject: deliverSubject,
		// FilterSubject:  filterSubject,
	}); err != nil {
		log.Print("Add Consumer "+channel+": Error: ", err)
	}
	return nil
}
