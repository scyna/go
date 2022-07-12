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
		Durable:      name,
		DeliverGroup: name,
		//DeliverSubject: deliverSubject,
		// FilterSubject:  filterSubject,
	}); err != nil {
		log.Print("Add Consumer "+channel+": Error: ", err)
	}
	return nil
}

func AddConsumer(stream string, durable string, group string, deliverSubject string, filterSubject string) error {

	if _, err := scyna.JetStream.AddConsumer(stream, &nats.ConsumerConfig{
		Durable:        durable,
		DeliverGroup:   group,
		DeliverSubject: deliverSubject,
		FilterSubject:  filterSubject,
	}); err != nil {
		log.Print("Add Consumer "+durable+": Error: ", err)
		return err
	}
	return nil
}
