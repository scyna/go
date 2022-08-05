package scyna

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2/qb"
)

const es_TRY_COUNT = 10
const es_BUCKET_SIZE = 1024

var err_ID_IS_RESERVED = errors.New("id is reserved")

var esBucket int64 = 1

type esStateType int

type EventStore struct {
	Bucket   int64
	ID       int64
	EntityID []uint64
	Time     time.Time
	Subject  string
	Data     []byte
}

const (
	ES_GET_LAST_ID     esStateType = 1
	ES_GET_LAST_BUCKET esStateType = 2
	ES_RESERVE_BUCKET  esStateType = 3
	ES_STORE_EVENT     esStateType = 4
)

func storeEvent(m *nats.Msg) (bool, int64) {
	tryCount := 0
	state := ES_GET_LAST_ID
	lastBucket := esBucket
	var lastID int64
	var err error

	for tryCount < es_TRY_COUNT {
		switch state {
		case ES_GET_LAST_ID:
			if lastID, err = getLastID(lastBucket); err == gocql.ErrNotFound {
				lastID = (lastBucket - 1) * es_BUCKET_SIZE
				state = ES_STORE_EVENT
				continue
			}
			if err != nil {
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
				state = ES_RESERVE_BUCKET
				continue
			}
			state = ES_GET_LAST_ID
		case ES_RESERVE_BUCKET:
			if err = reserveBucket(lastBucket); err == nil {
				state = ES_STORE_EVENT
				continue
			}
			tryCount++
		case ES_STORE_EVENT:
			nextID := lastID + 1
			if err := appendEvent(nextID, m); err == nil {
				return true, nextID
			}
			tryCount++
			if err == err_ID_IS_RESERVED {
				state = ES_GET_LAST_ID
				continue
			}
		}
	}
	return false, 0
}

func reserveBucket(bucket int64) error {
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
	if err := qb.Select(module + ".event_store").
		Columns("id").
		Where(qb.Eq("bucket")).
		Query(DB).Bind(0).
		GetRelease(&lastBucket); err != nil {
		return 0, err
	}
	return lastBucket, nil
}

func getLastID(bucket int64) (int64, error) {
	var lastID int64
	if err := qb.Select(module + ".event_store").
		Columns("id").
		Where(qb.Eq("bucket")).
		Query(DB).Bind(bucket).
		GetRelease(&lastID); err != nil {
		return 0, err
	}
	return lastID, nil
}

func appendEvent(id int64, m *nats.Msg) error {
	bucket := id/es_BUCKET_SIZE + 1
	if applied, err := qb.Insert(module+".event_store").
		Columns("bucket", "id", "subject", "data", "time").
		Unique().
		Query(DB).
		Bind(bucket, id, m.Subject, m.Data, time.Now()).
		ExecCASRelease(); !applied {
		if err != nil {
			return err
		}
		return err_ID_IS_RESERVED
	}
	esBucket = bucket
	return nil
}

func GetEvent(eventID int64) *EventStore {
	var eventStore EventStore
	bucket := eventID/es_BUCKET_SIZE + 1
	if err := qb.Select(module+".event_store").
		Columns("bucket", "id", "subject", "data", "time", "entity_id").
		Where(qb.Eq("bucket"), qb.Eq("id")).
		Query(DB).
		Bind(bucket, eventID).
		GetRelease(&eventStore); err != nil {
		return nil
	}
	return &eventStore
}
