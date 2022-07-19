package scyna

import (
	"errors"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
)

const es_TRY_COUNT = 10
const es_BUCKET_SIZE = 1024

var err_ID_IS_RESERVED = errors.New("id is reserved")

var esBucket int64 = 1

type esStateType int

const (
	ES_GET_LAST_ID     esStateType = 1
	ES_GET_LAST_BUCKET esStateType = 2
	ES_UPDATE_BUCKET   esStateType = 3
	ES_STORE_EVENT     esStateType = 4
)

func StoreEvent(m *nats.Msg) bool {
	tryCount := 0
	state := ES_GET_LAST_ID
	lastBucket := esBucket
	var lastID int64
	var err error

	for tryCount < es_TRY_COUNT {
		switch state {
		case ES_GET_LAST_ID:
			if lastID, err = getLastID(lastBucket); err != nil {
				tryCount++
				continue
			}
			if lastID == lastBucket*es_BUCKET_SIZE { /*reach end of bucket*/
				state = ES_GET_LAST_BUCKET
				continue
			}
			state = ES_STORE_EVENT
		case ES_GET_LAST_BUCKET:
			if lastBucket, err = getLastBucket(); err != nil {
				tryCount++
				continue
			}
			if lastBucket == esBucket { /*need to switch bucket*/
				lastBucket++
				state = ES_UPDATE_BUCKET
				continue
			}
			state = ES_STORE_EVENT
		case ES_UPDATE_BUCKET:
			if err = saveLastBucket(lastBucket); err == nil {
				state = ES_STORE_EVENT
				continue
			}
			tryCount++
		case ES_STORE_EVENT:
			if err := saveEventToStore(lastID+1, m); err == nil {
				esBucket = lastBucket
				return true
			}
			tryCount++
			if err == err_ID_IS_RESERVED {
				state = ES_GET_LAST_ID
				continue
			}
		}
	}
	return false
}

func saveLastBucket(bucket int64) error {
	if applied, err := qb.Insert(module+".event_store").
		Columns("bucket", "id").
		Unique().
		Query(DB).
		Bind(0, bucket).
		ExecCASRelease(); !applied {
		return err
	}
	return nil
}

func getLastBucket() (int64, error) {
	var lastBucket int64
	if err := DB.Session.Query("SELECT id FROM " + module + ".event_store WHERE bucket=0 LIMIT 1").
		Consistency(gocql.One).
		Scan(&lastBucket); err != nil {
		return 0, err
	}
	return lastBucket, nil
}

func getLastID(bucket int64) (int64, error) {
	var lastID int64
	if err := DB.Session.Query("SELECT id FROM "+module+".event_store WHERE bucket=? LIMIT 1", bucket).
		Consistency(gocql.One).
		Scan(&lastID); err != nil {
		return 0, err
	}
	return lastID, nil
}

func saveEventToStore(id int64, m *nats.Msg) error {
	bucket := id/es_BUCKET_SIZE + 1
	if applied, err := qb.Insert(module+".event_store").
		Columns("bucket", "id", "subject", "data").
		Unique().
		Query(DB).
		Bind(bucket, id, m.Subject, m.Data).
		ExecCASRelease(); !applied {
		if err != nil {
			return err
		}
		return err_ID_IS_RESERVED
	}
	return nil
}
