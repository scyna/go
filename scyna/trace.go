package scyna

import (
	"time"

	"github.com/gocql/gocql"
)

type TraceType uint32

const (
	TRACE_SERVICE TraceType = 1
	TRACE_EVENT   TraceType = 2
	TRACE_SIGNAL  TraceType = 3
	TRACE_SYNC    TraceType = 4
	TRACE_TASK    TraceType = 5
)

type Trace struct {
	ParentID    uint64    `db:"parent_id"`
	ID          uint64    `db:"id"`
	Type        TraceType `db:"type"`
	Time        time.Time `db:"time"`
	Duration    uint64    `db:"duration"`
	Path        string    `db:"path"`
	Source      string    `db:"source"`
	SessionID   uint64    `db:"session_id"`
	Status      int32     `db:"status"`
	RequestBody string
}

func (trace *Trace) Record() {
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	EmitSignal(TRACE_CREATED_CHANNEL, &TraceCreatedSignal{
		ID:        trace.ID,
		ParentID:  trace.ParentID,
		Type:      uint32(trace.Type),
		Time:      uint64(trace.Time.UnixMicro()),
		Duration:  trace.Duration,
		Path:      trace.Path,
		Source:    trace.Source,
		SessionID: trace.SessionID,
		Status:    trace.Status,
	})
}

func (trace *Trace) Save() {
	day := GetDayByTime(time.Now())
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	qBatch := DB.NewBatch(gocql.LoggedBatch)
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
	qBatch.Query("INSERT INTO scyna.app_has_trace(app_code, trace_id, day) VALUES (?,?,?)",
		trace.Source,
		trace.ID,
		day,
	)
	qBatch.Query("INSERT INTO scyna.tag(trace_id, key, value) VALUES (?,?,?)",
		trace.ID,
		"request",
		trace.RequestBody,
	)
	if err := DB.ExecuteBatch(qBatch); err != nil {
		LOG.Error("Can not save trace - " + err.Error())
	}
}
