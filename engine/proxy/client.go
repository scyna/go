package proxy

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type Client struct {
	ID     string `db:"id"`
	Secret string `db:"secret"`
}

func (proxy *Proxy) initClients() {
	proxy.Clients = proxy.loadClients()
	_, err := scyna.Connection.Subscribe(scyna.CLIENT_UPDATE_CHANNEL, func(msg *nats.Msg) {
		scyna.LOG.Info("Reload Clients")
		proxy.Clients = proxy.loadClients()
	})
	if err != nil {
		fmt.Println("initClients: " + err.Error())
	}
}

func (proxy *Proxy) loadClients() map[string]Client {
	ret := make(map[string]Client)
	var clients []Client

	if err := qb.Select("scyna.client").
		Columns("id", "secret").
		Query(scyna.DB).
		SelectRelease(&clients); err == nil {
		for _, c := range clients {
			ret[c.ID] = c
		}
	} else {
		scyna.LOG.Error("Load Clients fail: " + err.Error())
	}
	return ret
}
