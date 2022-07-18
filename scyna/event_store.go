package scyna

import (
	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
)

const es_TRY_COUNT = 10

var lastEventID int64

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
	/*load event with id = 0, data hold lastID of event */
	if err := qb.Select(module + ".event_store").
		Columns("blobAsBigint(data)").
		Where(qb.Eq("id")).
		Query(DB).Bind(0).
		Get(&lastEventID); err != nil {
		return err
	}
	return nil
}

func saveEventToStore(m *nats.Msg) error {
	batch := DB.NewBatch(gocql.LoggedBatch)
	nextID := lastEventID + 1

	batch.Query("INSERT INTO "+module+".event_store(id, subject, data) VALUES(?,?,?) IF NOT EXISTS", nextID, m.Subject, m.Data)
	batch.Query("UPDATE "+module+".event_store SET data=bigintAsBlob(?) WHERE id=?", nextID, 0)

	if applied, _, err := DB.ExecuteBatchCAS(batch); applied {
		return nil
	} else {
		return err
	}
}
