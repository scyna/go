package scyna

import (
	"fmt"
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
)

type eventStore struct {
	version uint64
	esQuery string
}

var EventStore *eventStore

func InitEventStore(keyspace string) {
	var version uint64 = 0

	if err := qb.Select(keyspace + ".event_store").
		Max("event_id").
		Query(DB).
		GetRelease(&version); err != nil {
		log.Fatal("Can not init EventStore")
	}

	/*TODO: push last event*/

	EventStore = &eventStore{
		version: version,
		esQuery: fmt.Sprintf("INSERT INTO %s.event_store(event_id, entity_id, channel, data) VALUES(?,?,?,?)", keyspace),
	}
}

func (es *eventStore) Add(ctx *Command, aggregate uint64, channel string, event proto.Message) bool {
	var id = es.version + 1

	bytes, err := proto.Marshal(event)
	if err != nil {
		ctx.Logger.Error("Can not marshal event data")
		return false
	}

	ctx.Batch.Query(es.esQuery, id, aggregate, channel, bytes)

	if err := DB.ExecuteBatch(ctx.Batch); err == nil {
		es.version = id
		return true
	}

	return false
}
