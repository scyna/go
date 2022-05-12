package proxy

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

func (proxy *Proxy) saveTrace(trace *scyna.Trace) {
	day := scyna.GetDayByTime(time.Now())
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.trace(type, path, day, id, time, duration, session_id, source, status) VALUES (?,?,?,?,?,?,?,?,?)",
		trace.Type,
		trace.Path,
		day,
		trace.ID,
		trace.Time,
		trace.Duration,
		trace.SessionID,
		trace.Source,
		trace.Status,
	)
	qBatch.Query("INSERT INTO scyna.client_has_trace(client_id, trace_id, day) VALUES (?,?,?)",
		trace.Source,
		trace.ID,
		day,
	)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		scyna.LOG.Error("Can not save trace - " + err.Error())
	}
}
