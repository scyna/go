package trace

import (
	"log"
	"strconv"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func TraceCreated(signal *scyna.TraceCreatedSignal) {
	day := scyna.GetDayByTime(time.Now())
	var source *string = nil
	if len(signal.Source) > 0 {
		source = &signal.Source
	}

	if signal.ParentID == 0 {
		if err := qb.Insert("scyna.trace").
			Columns("type", "path", "day", "id", "time", "duration", "session_id", "source", "status").
			Query(scyna.DB).
			Bind(
				signal.Type,
				signal.Path,
				day,
				signal.ID,
				time.UnixMicro(int64(signal.Time)),
				signal.Duration,
				signal.SessionID,
				source,
				signal.Status).
			ExecRelease(); err != nil {
			log.Print("Can not save trace created " + strconv.FormatUint(signal.ID, 10) + " / " + err.Error())
		}
	} else {
		qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
		qBatch.Query("INSERT INTO scyna.trace(type, path, day, id, time, duration, session_id, parent_id, source, status)"+
			" VALUES (?,?,?,?,?,?,?,?,?,?)",
			signal.Type,
			signal.Path,
			day,
			signal.ID,
			time.UnixMicro(int64(signal.Time)),
			signal.Duration,
			signal.SessionID,
			signal.ParentID,
			source,
			signal.Status)
		qBatch.Query("INSERT INTO scyna.span(parent_id, child_id) VALUES (?,?)", signal.ParentID, signal.ID)

		if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
			log.Print("Can not save trace created " + strconv.FormatUint(signal.ID, 10) + " / " + strconv.FormatUint(signal.ParentID, 10) + " / " + err.Error())
		}
	}

}
