package proxy

import (
	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
	"time"
)

func (proxy *Proxy) saveContext(ctx *scyna.Context) {
	day := scyna.GetDayByTime(time.Now())
	ctx.Duration = uint64(time.Now().UnixNano() - ctx.Time.UnixNano())
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.trace(type, path, day, id, time, duration, session_id, source, status) "+
		" VALUES (?,?,?,?,?,?,?,?,?)", ctx.Type, ctx.Path, day, ctx.ID, ctx.Time, ctx.Duration, ctx.SessionID, ctx.Source, ctx.Status)
	qBatch.Query("INSERT INTO scyna.client_has_trace(client_id, trace_id, day) VALUES (?,?,?)", ctx.Source, ctx.ID, day)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		scyna.LOG.Error("Can not save trace - " + err.Error())
	}
}
