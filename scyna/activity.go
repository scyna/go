package scyna

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

var activityPool *QueryPool

func SetupActivity(table string) {
	activityPool = NewQueryPool(func() *gocqlx.Queryx {
		return qb.Insert(module+".activity").
			Columns("entity_id", "event_id").
			Unique().
			Query(scyna.DB)
	})
}

func addActivity(entityID uint64, eventID int64) {
	qInsert := activityPool.GetQuery()

	if applied, err := qInsert.Bind(entityID, eventID).ExecCAS(); applied {
		return
	} else if err != nil {
		LOG.Error("addActivity" + err.Error())
		return
	}
}
