package scyna

import (
	"time"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

const tryCount = 10

type ActivityStream struct {
	keyspace string
	queries  *QueryPool
}

type Activity struct {
	EntityID uint64    `db:"entity_id"`
	Type     int32     `db:"type"`
	Time     time.Time `db:"time"`
	Data     []byte    `db:"data"`
}

func InitActivityStream(keyspace string) *ActivityStream {
	return &ActivityStream{
		keyspace: keyspace,
		queries: NewQueryPool(func() *gocqlx.Queryx {
			return qb.Insert(keyspace+".activity").
				Columns("entity_id", "type", "time", "data").
				Unique().
				Query(DB)
		}),
	}
}

func (stream *ActivityStream) Add(entity uint64, Type int32, activity protoreflect.ProtoMessage) {
	t := uint64(time.Now().UnixMicro())

	var data []byte
	if activity != nil {
		data, _ = proto.Marshal(activity)
	}

	qInsert := stream.queries.GetQuery()
	defer stream.queries.Put(qInsert)

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

func (stream *ActivityStream) List(entity uint64) []Activity {
	var ret []Activity
	if err := qb.Select(stream.keyspace+".activity").
		Columns("entity_id", "type", "time", "data").
		Where(qb.Eq("entity_id")).
		Query(DB).
		Bind(entity).
		SelectRelease(&ret); err != nil {
		LOG.Error("Can not get event: " + err.Error())
	}
	return ret
}
