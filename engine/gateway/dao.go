package gateway

import (
	"log"

	"github.com/scyna/go/scyna"

	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
)

func (gateway *Gateway) saveCall(app string, id uint64, day int, start time.Time, duration int64, url string, caller string, context *scyna.HttpContext) {
	// qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	// qBatch.Query("INSERT INTO scyna.call(id, day, time, duration, source,request, response, status, session_id, caller_id)"+
	// 	" VALUES (?,?,?,?,?,?,?,?,?,?)", id, day, start, duration, url, service.Request.Body, service.Response.Body,
	// 	service.Response.Code, service.Response.SessionID, caller,
	// )
	// qBatch.Query("INSERT INTO scyna.app_has_call(app_code, call_id, day) VALUES (?,?,?)", app, id, day)
	// if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
	// 	log.Print(err)
	// 	scyna.LOG.Error("Error in save call")
	// }
}

func (gateway *Gateway) saveErrorCall(app string, status int, id uint64, day int, start time.Time, url string, caller string) {
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.call(id, day, time, source, status, caller_id)"+
		" VALUES (?,?,?,?,?,?)", id, day, start, url, status, caller)
	qBatch.Query("INSERT INTO scyna.app_has_call(app_code, call_id, day) VALUES (?,?,?)", app, id, day)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		log.Print(err)
		scyna.LOG.Error("Error in save call")
	}
}

func updateSesion(token string, exp time.Time) bool {
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
