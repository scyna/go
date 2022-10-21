package scyna

import "google.golang.org/protobuf/proto"

type eventStore struct {
	Version uint64
}

func NewEventStore() *eventStore {
	/*TODO: load lastest version from database*/
	return &eventStore{
		Version: 0, /*FIXME*/
	}
}

func (es *eventStore) Add(ctx *Command, aggregate uint64, channel string, event proto.Message) bool {
	/*TODO: add event to EventStore (batch)*/
	ctx.batch.Query("TODO")

	if err := DB.ExecuteBatch(ctx.batch); err == nil {
		return true
	}
	return false
}
