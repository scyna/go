package scyna

import (
	"sync"
	"time"

	"github.com/scylladb/gocqlx/v2"

	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

const tryCount = 10

type ActivityStream struct {
	keyspace string
	Queries  *QueryPool
}

type Activity struct {
	EntityID uint64    `db:"entity_id"`
	Type     int32     `db:"type"`
	Time     time.Time `db:"time"`
	Data     []byte    `db:"data"`
}

func InitActivityStream(keyspace string) *ActivityStream {
	tName := keyspace + ".activity"
	return &ActivityStream{
		keyspace: keyspace,
		Queries: &QueryPool{
			sync.Pool{
				New: func() interface{} {
					return qb.Insert(tName).Columns("entity_id", "type", "time", "data").Unique().Query(DB)
				},
			},
		},
	}
}

func (stream *ActivityStream) Add(entity uint64, Type int, activity protoreflect.ProtoMessage) {
	t := uint64(time.Now().UnixMicro())

	var data []byte
	if activity != nil {
		data, _ = proto.Marshal(activity)
	}

	qInsert := stream.Queries.GetQuery()
	defer stream.Queries.Put(qInsert)

	for i := 0; i < tryCount; i++ {
		qInsert.Bind(entity, Type, t, data)
		if applied, err := qInsert.ExecCAS(); applied {
			return
		} else {
			if err != nil {
				LOG.Error("ActivityStream.Add :" + err.Error())
				return
			}
		}
		t++
	}
}

func GetActivityQuery(name string, entity uint64) *gocqlx.Queryx {
	return qb.Select(name).
		Columns("entity_id", "type", "time", "data").
		Where(qb.Eq("entity_id")).
		Query(DB).
		Bind(entity)
}

func (stream *ActivityStream) List(entity uint64) []Activity {
	tName := stream.keyspace + ".activity"
	qSelect := GetActivityQuery(tName, entity)
	var event []Activity
	if err := qSelect.Select(&event); err != nil {
		LOG.Error("Can not get event: " + err.Error())
	}

	return event
}
