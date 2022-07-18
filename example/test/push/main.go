package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect("localhost:4222")

	js, _ := nc.JetStream()

	js.Subscribe("account.created", func(m *nats.Msg) {
		m.Ack()
		println("created:", string(m.Data))
	}, nats.Durable("c1"), nats.ManualAck())

	js.Subscribe("account.updated", func(m *nats.Msg) {
		m.Ack()
		println("updated:", string(m.Data))
	}, nats.Durable("c1"), nats.ManualAck())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
