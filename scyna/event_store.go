package scyna

import (
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
)

const es_TRY_COUNT = 10
const es_BUCKET_SIZE = 1024

var errCanNotApplied = errors.New("Can not save bucket")

var esBucket int64

func InitEventStore() {
	getLastBucket()
}

func storeEvent(m *nats.Msg) bool {
	for i := 0; i < es_TRY_COUNT; i++ {
		if lastId, err := getLastEventID(esBucket); err == nil {
			if saveEventToStore(lastId+1, m) == nil {
				return true
			}
		}
	}
	return false
}

func getLastBucket() {
	/*TODO*/
}

func getLastEventID(bucket int64) (int64, error) {
	/*TODO*/
	return 0, nil
}

func saveEventToStore(id int64, m *nats.Msg) error {
	bucket := esBucket
	if id > esBucket*es_BUCKET_SIZE {
		bucket++
	}

	if applied, err := qb.Insert(module+".event_store").
		Columns("bucket", "id", "subject", "data").
		Unique().
		Query(DB).
		Bind(bucket, id, m.Subject, m.Data).
		ExecCASRelease(); !applied {

		if err != nil {
			return err
		}

		if bucket > esBucket {
			getLastBucket()
		}
		return errCanNotApplied
	}

	if bucket > esBucket {
		/*TODO:update bucket*/
		esBucket = bucket
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
