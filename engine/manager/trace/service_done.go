package trace

import (
	"fmt"
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func ServiceDone(signal *scyna.ServiceDoneSignal) {
	log.Print("Service Done")
	data := fmt.Sprintf("{Status:%d}\n%s\n%s", signal.Status, signal.Request, signal.Response)
	log.Print(data)
	qb.Insert("scyna.tag").
		Columns("trace_id", "key", "value").
		Query(scyna.DB).
		Bind(signal.TraceID, "data", data).
		ExecRelease()
}
