package scyna

import (
	"time"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

const tryCount = 10

var activityPool *QueryPool

func SetupActivity(table string) {
	activityPool = NewQueryPool(func() *gocqlx.Queryx {
		return qb.Insert(table).
			Columns("entity_id", "type", "time", "data").
			Unique().
			Query(DB)
	})
}

func AddActivity(entity uint64, Type int32, activity protoreflect.ProtoMessage) {
	if activityPool == nil {
		LOG.Error("Setup activity table first")
		return
	}

	t := time.Now()

	var data []byte
	if activity != nil {
		data, _ = proto.Marshal(activity)
	}

	qInsert := activityPool.GetQuery()
	defer activityPool.Put(qInsert)

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
		t.Add(time.Millisecond + 1)
	}
}
