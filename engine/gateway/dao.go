package gateway

import (
	"log"

	"github.com/gocql/gocql"

	"github.com/scyna/go/scyna"

	"time"

	"github.com/scylladb/gocqlx/v2/qb"
)

func (gateway *Gateway) saveContext(trace *scyna.Trace) {
	day := scyna.GetDayByTime(time.Now())
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.trace(type, path, day, id, time, duration, session_id, source, status) "+
		" VALUES (?,?,?,?,?,?,?,?,?)", trace.Type, trace.Path, day, trace.ID, trace.Time, trace.Duration, trace.SessionID, trace.Source, trace.Status)
	qBatch.Query("INSERT INTO scyna.app_has_trace(app_code, trace_id, day) VALUES (?,?,?)", trace.Source, trace.ID, day)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		scyna.LOG.Error("Can not save trace - " + err.Error())
	}
}

func updateSession(token string, exp time.Time) bool {
	err := qb.Update("scyna.authentication").
		Set("expired").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(exp, token).
		ExecRelease()
	return err == nil
}

func checkService(token string, app string, url string) *time.Time {
	/*check authentication*/
	var auth struct {
		Expired time.Time `db:"expired"`
		Apps    []string  `db:"apps"`
	}

	if err := qb.Select("scyna.authentication").
		Columns("expired", "apps").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(token).
		GetRelease(&auth); err != nil {
		log.Println("authentication", err.Error())
		return nil
	}

	hasApp := false
	for _, a := range auth.Apps {
		if a == app {
			hasApp = true
			break
		}
	}

	if !hasApp {
		log.Print("No app")
		return nil
	}

	// /*check app_use_service*/
	// if err := qb.Select("scyna.app_use_service").
	// 	Columns("app_code").
	// 	Where(qb.Eq("app_code"), qb.Eq("service_url")).
	// 	Limit(1).
	// 	Query(scyna.DB).
	// 	Bind(app, url).
	// 	GetRelease(&app); err != nil {
	// 	log.Println("app_use_service",err.Error())
	// 	return false
	// }
	ret := auth.Expired
	return &ret
}
