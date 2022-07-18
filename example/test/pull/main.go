package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect("localhost:4222")

	js, _ := nc.JetStream()

	sub, err := js.PullSubscribe("", "c1", nats.BindStream("stream1"))
	if err != nil {
		panic(err)
	}

	for {
		msgs, _ := sub.Fetch(5)

		if len(msgs) > 0 {
			for _, msg := range msgs {
				msg.Ack()
				println(string(msg.Data))
			}
		}
	}
}
