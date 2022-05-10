package trace

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func WriteTrace(signal *scyna.TraceCreatedSignal) {
	log.Print("Write Trace to Database")
	day := scyna.GetDayByTime(time.Now())

	if signal.ParentID == 0 {
		if err := qb.Insert("scyna.trace").
			Columns("path", "day", "id", "time", "duration", "session_id").
			Query(scyna.DB).
			Bind(
				signal.Path,
				day,
				signal.ID,
				time.UnixMicro(int64(signal.Time)),
				signal.Duration,
				signal.SessionID).
			ExecRelease(); err != nil {
			log.Print(err)
		}
	} else {
		qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
		qBatch.Query("INSERT INTO scyna.trace(path, day, id, time, duration, session_id, parent_id)"+
			" VALUES (?,?,?,?,?,?,?)",
			signal.Path,
			signal.ID,
			day,
			time.UnixMicro(int64(signal.Time)),
			signal.Duration,
			signal.SessionID,
			signal.ParentID)
		qBatch.Query("INSERT INTO scyna.span(parent_id, child_id) VALUES (?,?)",
			signal.ParentID, signal.ID)

		if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
			log.Print(err)
		}
	}

}
