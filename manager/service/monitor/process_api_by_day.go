package monitor

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
	"log"
	"strings"
	"time"
)

func ProcessMonitorByDay(s *scyna.Service, request *proto.ProcessMonitorByDayRequest) {
	s.Logger.Info("Receive ProcessMonitorByDay request")
	if validateProcessMonitorByDay(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	date, err := time.Parse("2006-01-02", request.Day)
	if err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	day := scyna.GetDayByTime(date)

	var listAPI []distinct
	if err := qb.Select("scyna.trace").
		Distinct("day", "path").
		Where(qb.Eq("day")).
		AllowFiltering().
		Query(scyna.DB).
		Bind(day).
		SelectRelease(&listAPI); err != nil {
		s.Logger.Info(err.Error())
		s.Error(scyna.SERVER_ERROR)
		return
	}
	var totalTrace []trace
	for _, api := range listAPI {
		if !strings.HasPrefix(api.Path, "/") {
			continue
		}
		var t []trace
		if err := qb.Select("scyna.trace").
			Columns("path", "day", "id", "duration", "parent_id", "session_id",
				"source", "status", "time", "type").
			Where(qb.Eq("day"), qb.Eq("path")).
			Query(scyna.DB).
			Bind(api.Day, api.Path).
			SelectRelease(&t); err != nil {
			s.Logger.Info(api.Path + ": " + err.Error())
			s.Error(scyna.SERVER_ERROR)
			return
		}

		totalTrace = append(totalTrace, t...)
	}

	s.Done(scyna.OK)

	var slots []slot

	for i := 0; i < 24; i++ {
		slots = append(slots, slot{
			Time:    i,
			Success: 0,
			Error:   0,
		})
	}

	totalError := 0
	totalSuccess := 0
	totalPermission := 0
	avgLatency := 0

	minLatency := 0
	maxLatency := 0
	totalLatency := 0
	var tracePermission []trace
	var traceError []trace
	for i, t := range totalTrace {

		totalLatency += t.Duration
		if i == 0 || maxLatency < t.Duration {
			maxLatency = t.Duration
		}
		if i == 0 || minLatency > t.Duration {
			minLatency = t.Duration
		}
		if t.Status >= 500 {
			totalError = totalError + 1
			slots[t.Time.Hour()].Error = slots[t.Time.Hour()].Error + 1
			traceError = append(traceError, t)
		} else if t.Status == 401 {
			totalPermission = totalPermission + 1
			slots[t.Time.Hour()].Success = slots[t.Time.Hour()].Success + 1
			tracePermission = append(tracePermission, t)
		} else {
			totalSuccess = totalSuccess + 1
			slots[t.Time.Hour()].Success = slots[t.Time.Hour()].Success + 1
		}
	}

	avgLatency = totalLatency / len(totalTrace)

	log.Println(len(totalTrace))
	log.Println(totalLatency)
	log.Println(avgLatency)
	data, err := json.Marshal(slots)
	if err != nil {
		s.Logger.Info(err.Error())
		return
	}

	if err := qb.Insert("scyna.api_report_by_day").
		Columns("day", "total_error", "total_success", "avg_latency",
			"min_latency", "max_latency", "total_permission", "data").
		Query(scyna.DB).
		Bind(date, totalError, totalSuccess, avgLatency, minLatency, maxLatency, totalPermission, data).
		ExecRelease(); err != nil {
		s.Logger.Info(err.Error())
	}
	scynaSession := scyna.DB.Session
	batch := scynaSession.NewBatch(gocql.LoggedBatch)
	for _, t := range tracePermission {
		batch.Query("INSERT INTO scyna.api_report_by_permission(day,trace_id,path,client_id) VALUES (?,?,?,?);",
			date, t.Id, t.Path, t.Source)
	}
	for _, t := range traceError {
		batch.Query("INSERT INTO scyna.api_report_by_error(day,trace_id,path,client_id) VALUES (?,?,?,?);",
			date, t.Id, t.Path, t.Source)
	}

	if err := scynaSession.ExecuteBatch(batch); err != nil {
		s.Logger.Info("Save batch error: " + err.Error())
	}
}

type slot struct {
	Time    int `json:"time"`
	Error   int `json:"error"`
	Success int `json:"success"`
}

type trace struct {
	Path      string    `db:"path"`
	Day       int       `db:"day"`
	Id        int64     `db:"id"`
	Duration  int       `db:"duration"`
	ParentId  string    `db:"parent_id"`
	SessionId int64     `db:"session_id"`
	Source    string    `db:"source"`
	Status    int       `db:"status"`
	Time      time.Time `db:"time"`
	Type      int       `db:"type"`
}

type distinct struct {
	Day  int32  `db:"day"`
	Path string `db:"path"`
}

func validateProcessMonitorByDay(request *proto.ProcessMonitorByDayRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Day, validation.Required, validation.Length(5, 100)))
}
