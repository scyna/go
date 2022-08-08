package scyna

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

var activityPool *QueryPool

func SetupActivity(table string) {
	activityPool = NewQueryPool(func() *gocqlx.Queryx {
		return qb.Insert(table).
			Columns("entity_id", "event_id").
			Unique().
			Query(DB)
	})
}

func addActivity(entityID uint64, eventID int64) {
	if activityPool == nil {
		LOG.Error("Setup activity table first")
		return
	}

	qInsert := activityPool.GetQuery()
	defer activityPool.Put(qInsert)

	qInsert.Bind(entityID, eventID)
	if applied, err := qInsert.ExecCAS(); applied {
		return
	} else if err != nil {
		LOG.Error("ActivityStream.Add :" + err.Error())
		return
	}
}
