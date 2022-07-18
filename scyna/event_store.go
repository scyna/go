package scyna

import (
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
)

const es_TRY_COUNT = 10
const es_BUCKET_SIZE = 1024

var errBucketUpdated = errors.New("bucket updated")
var errIDIsReserved = errors.New("id is reserved")

var esBucket int64

func InitEventStore() {
	esBucket, _ = getLastBucket()
}

func StoreEvent(m *nats.Msg) bool {
	tryCount := 0
	state := "GETID"
	lastBucket := esBucket
	var lastID int64
	var err error

	for tryCount < es_BUCKET_SIZE {
		switch state {
		case "GET_ID":
			if lastID, err = getLastEventID(lastBucket); err != nil {
				tryCount++
				continue
			}
			if lastID == esBucket*es_BUCKET_SIZE { /*reach end of bucket*/
				lastBucket++
				state = "SAVE_BUCKET"
				continue
			}
			state = "SAVE_EVENT"
		case "GET_BUCKET":
			if lastBucket, err = getLastBucket(); err != nil {
				tryCount++
				continue
			}
			if lastBucket == esBucket { /*need to switch bucket*/
				lastBucket++
				state = "SAVE_BUCKET"
				continue
			}
			state = "SAVE_EVENT"
		case "SAVE_BUCKET":
			if err = saveLastBucket(lastBucket); err == nil {
				state = "SAVE_EVENT"
				continue
			}
			if err == errBucketUpdated {
				state = "GET_ID"
				continue
			}
			tryCount++
		case "SAVE_EVENT":
			if err := saveEventToStore(lastID+1, m); err == nil {
				esBucket = lastBucket
				return true
			}
			tryCount++
			if err == errIDIsReserved {
				state = "GET_ID"
				continue
			}
		}
	}
	return false
}

func saveLastBucket(bucket int64) error {
	return nil
}

func getLastBucket() (int64, error) {
	/*TODO*/
	return 0, nil
}

func getLastEventID(bucket int64) (int64, error) {
	/*TODO*/
	return 0, nil
}

func saveEventToStore(id int64, m *nats.Msg) error {
	if applied, err := qb.Insert(module+".event_store").
		Columns("bucket", "id", "subject", "data").
		Unique().
		Query(DB).
		Bind(esBucket, id, m.Subject, m.Data).
		ExecCASRelease(); !applied {
		if err != nil {
			return err
		}
		return errIDIsReserved
	}
	return nil
}

// func getLastEventID() (int64, error) {
// 	/*load event with id = 0, data hold lastID of event */
// 	var lastEventID int64
// 	ctx := context.Background()
// 	if err := DB.Session.Query("SELECT blobAsBigint(data) as last_id FROM " + module + ".event_store WHERE id=0 LIMIT 1").
// 		WithContext(ctx).
// 		Consistency(gocql.One).
// 		Scan(&lastEventID); err != nil {
// 		return 0, err
// 	}
// 	return lastEventID, nil
// }

// func saveEventToStore(id int64, m *nats.Msg) error {
// 	batch := DB.NewBatch(gocql.LoggedBatch)
// 	batch.Query("INSERT INTO "+module+".event_store(id, subject, data) VALUES(?,?,?) IF NOT EXISTS", id, m.Subject, m.Data)
// 	batch.Query("UPDATE "+module+".event_store SET data=bigintAsBlob(?) WHERE id=?", id, 0)

// 	if applied, _, err := DB.ExecuteBatchCAS(batch); applied {
// 		return nil
// 	} else {
// 		return err
// 	}
// }

// func getLastBucket() (int64, error) {
// 	/*load event with id = 0, data hold lastID of event */
// 	var lastEventID int64
// 	ctx := context.Background()
// 	if err := DB.Session.Query("SELECT blobAsBigint(data) as last_id FROM " + module + ".event_store WHERE id=0 LIMIT 1").
// 		WithContext(ctx).
// 		Consistency(gocql.One).
// 		Scan(&lastEventID); err != nil {
// 		return 0, err
// 	}
// 	return lastEventID, nil
// }

// func getLastID(bucket int64) (int64, error) {
// 	/*load event with id = 0, data hold lastID of event */
// 	var lastEventID int64
// 	ctx := context.Background()
// 	if err := DB.Session.Query("SELECT blobAsBigint(data) as last_id FROM " + module + ".event_store WHERE id=0 LIMIT 1").
// 		WithContext(ctx).
// 		Consistency(gocql.One).
// 		Scan(&lastEventID); err != nil {
// 		return 0, err
// 	}
// 	return lastEventID, nil
// }
