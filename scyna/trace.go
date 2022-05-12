package scyna

import (
	"time"
)

type TraceType uint32

const (
	TRACE_SERVICE TraceType = 1
	TRACE_EVENT   TraceType = 2
	TRACE_SIGNAL  TraceType = 3
)

type Trace struct {
	ParentID  uint64    `db:"parent_id"`
	ID        uint64    `db:"id"`
	Type      TraceType `db:"type"`
	Time      time.Time `db:"time"`
	Duration  uint64    `db:"duration"`
	Path      string    `db:"path"`
	Source    string    `db:"source"`
	SessionID uint64    `db:"session_id"`
	Status    int32     `db:"status"`
}

func (trace *Trace) Save() {
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	EmitSignalLite(TRACE_CREATED_CHANNEL, &TraceCreatedSignal{
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
