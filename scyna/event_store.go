package scyna

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
)

const es_TRY_COUNT = 10

func storeEvent(m *nats.Msg) bool {
	for i := 0; i < es_TRY_COUNT; i++ {
		if err, lastID := getLastEventID(); err == nil {
			if saveEventToStore(lastID+1, m) == nil {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func getLastEventID() (error, int64) {
	/*load event with id = 0, data hold lastID of event */
	var lastEventID int64
	ctx := context.Background()
	if err := DB.Session.Query("SELECT blobAsBigint(data) as last_id FROM " + module + ".event_store WHERE id=0 LIMIT 1").
		WithContext(ctx).
		Consistency(gocql.One).
		Scan(&lastEventID); err != nil {
		return err, 0
	}
	return nil, lastEventID
}

func saveEventToStore(id int64, m *nats.Msg) error {
	batch := DB.NewBatch(gocql.LoggedBatch)
	batch.Query("INSERT INTO "+module+".event_store(id, subject, data) VALUES(?,?,?) IF NOT EXISTS", id, m.Subject, m.Data)
	batch.Query("UPDATE "+module+".event_store SET data=bigintAsBlob(?) WHERE id=?", id, 0)

	if applied, _, err := DB.ExecuteBatchCAS(batch); applied {
		return nil
	} else {
		return err
	}
}
