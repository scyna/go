package proxy

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func (proxy *Proxy) saveContext(trace *scyna.Trace) {
	trace.Save() //FIXME: direct save
	if len(trace.Source) > 0 {
		if err := qb.Insert("scyna.client_has_trace").
			Columns("client_id", "trace_id").
			Unique().
			Query(scyna.DB).
			Bind(&trace.Source, trace.ID).
			ExecRelease(); err != nil {
			scyna.LOG.Info("Can not save app_has_trace bc " + err.Error())
		}
	}
}
