package scyna

import (
	"fmt"
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
)

type eventStore struct {
	version       uint64
	esQuery       string
	activityQuery string
}

func InitEventStore(name string) *eventStore {
	var version uint64 = 0

	esTable := fmt.Sprintf("%s.%s_event_store", context, name)

	if err := qb.Select(esTable).
		Max("event_id").
		Query(DB).
		GetRelease(&version); err != nil {
		log.Fatal("Can not init EventStore")
	}

	/*TODO: push last event*/

	return &eventStore{
		version:       version,
		esQuery:       fmt.Sprintf("INSERT INTO %s.%s_event_store(event_id, aggregate_id, channel, data) VALUES(?,?,?,?)", context, name),
		activityQuery: fmt.Sprintf("INSERT INTO %s.%s_activity(aggregate_id, event_id) VALUES(?,?)", context, name),
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
	ctx.Batch.Query(es.activityQuery, aggregate, id)

	if err := DB.ExecuteBatch(ctx.Batch); err == nil {
		es.version = id
		return true
	}

	return false
}
