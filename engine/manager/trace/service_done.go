package trace

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func ServiceDone(signal *scyna.ServiceDoneSignal) {
	qb.Insert("scyna.tag").
		Columns("trace_id", "key", "value").
		Query(scyna.DB).
		Bind(signal.TraceID, "data", signal.Request+"\n"+signal.Response).
		ExecRelease()
}
