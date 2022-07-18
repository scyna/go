package scyna

import (
	"strconv"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
)

const es_TRY_COUNT = 10

type eventInStore struct {
	ID      int64
	Subject string
	Data    []byte
}

var eventStoreHeader eventInStore

func storeEvent(m *nats.Msg) bool {
	for i := 0; i < es_TRY_COUNT; i++ {
		if err := loadEventStoreHeader(); err == nil {
			if saveEventToStore(m) == nil {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func loadEventStoreHeader() error {
	/*load event with id = 0, Subject hold last event id */
	if err := qb.Select(module+".event_store").
		Columns("id", "subject").
		Where(qb.Eq("id")).
		Query(DB).Bind(0).
		Get(&eventStoreHeader); err != nil {
		return err
	}
	return nil
}

func saveEventToStore(m *nats.Msg) error {
	batch := DB.NewBatch(gocql.LoggedBatch)
	var nextID int64
	if id, err := strconv.ParseInt(eventStoreHeader.Subject, 10, 64); err != nil {
		return err
	} else {
		nextID = id + 1
	}

	batch.Query("INSERT INTO "+module+".event_store(id, subject, data) VALUES(?,?,?) IF NOT EXISTS", nextID, m.Subject, m.Data)
	batch.Query("UPDATE "+module+".event_store SET subject=? WHERE id=?", strconv.FormatInt(nextID, 10), 0) /*update last event id to header*/

	if applied, _, err := DB.ExecuteBatchCAS(batch); applied {
		return nil
	} else {
		return err
	}
}
