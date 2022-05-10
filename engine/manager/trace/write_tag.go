package trace

import (
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func WriteTag(signal *scyna.TagCreatedSignal) {
	log.Print("Write Tag to Database")

	qb.Insert("scyna.tag").
		Columns("trace_id", "key", "value").
		Query(scyna.DB).
		Bind(signal.TraceID, signal.Key, signal.Value).
		ExecRelease()
}
