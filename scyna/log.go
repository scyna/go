package scyna

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
)

type LogLevel int

const (
	LOG_INFO    LogLevel = 1
	LOG_ERROR   LogLevel = 2
	LOG_WARNING LogLevel = 3
	LOG_DEBUG   LogLevel = 4
	LOG_FATAL   LogLevel = 5
)

type LogData struct {
	Level    LogLevel
	Message  string
	ID       uint64
	Sequence uint64
	Session  bool
}

type Logger struct {
	session bool
	ID      uint64
}

var logQueue chan LogData

func UseDirectLog(count int) {
	logQueue = make(chan LogData)

	for i := 0; i < count; i++ {
		go func() {
			qSession := qb.Insert("scyna.session_log").Columns("session_id", "day", "time", "seq", "level", "message").Unique().Query(DB)
			qService := qb.Insert("scyna.log").Columns("trace_id", "time", "seq", "level", "message").Unique().Query(DB)
			for l := range logQueue {
				time_ := time.Now()
				if l.Session {
					if _, err := qSession.Bind(l.ID, GetDayByTime(time_), time_, l.Sequence, l.Level, l.Message).
						ExecCAS(); err != nil {
						log.Println("saveSessionLog: " + err.Error())
					}
				} else {
					if _, err := qService.Bind(l.ID, time_, l.Sequence, l.Level, l.Message).
						ExecCAS(); err != nil {
						log.Println("saveServiceLog: " + err.Error())
					}
				}
			}
		}()
	}
}

func UseRemoteLog(count int) {
	logQueue = make(chan LogData)

	for i := 0; i < count; i++ {
		go func() {
			for l := range logQueue {
				time_ := time.Now().UnixMicro()
				event := LogCreatedSignal{
					Time:    uint64(time_),
					ID:      l.ID,
					Level:   uint32(l.Level),
					Text:    l.Message,
					Session: l.Session,
					SEQ:     l.Sequence,
				}
				EmitSignal(LOG_CREATED_CHANNEL, &event)
			}
		}()
	}
}

func AddLog(data LogData) {
	if logQueue != nil {
		logQueue <- data
	}
}

func releaseLog() {
	if logQueue != nil {
		close(logQueue)
	}
}

func (l *Logger) writeLog(level LogLevel, message string) {
	message = formatLog(message)
	log.Print(message) //FIXME: for debug only
	if l.ID > 0 {
		AddLog(LogData{
			ID:       l.ID,
			Sequence: Session.NextSequence(),
			Level:    level,
			Message:  message,
			Session:  l.session,
		})
	}
}

func (l *Logger) Reset(id uint64) {
	l.ID = id
}

func (l *Logger) Info(messsage string) {
	l.writeLog(LOG_INFO, messsage)
}

func (l *Logger) Error(messsage string) {
	l.writeLog(LOG_ERROR, messsage)
}

func (l *Logger) Warning(messsage string) {
	l.writeLog(LOG_WARNING, messsage)
}

func (l *Logger) Debug(messsage string) {
	l.writeLog(LOG_DEBUG, messsage)
}

func (l *Logger) Fatal(messsage string) {
	l.writeLog(LOG_FATAL, messsage)
}

func formatLog(message string) string {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return fmt.Sprintf("[?:0 - ?] %s", message)
	}
	path := strings.Split(file, "/")
	filename := path[len(path)-1]

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return fmt.Sprintf("[%s:%d - ?] %s", filename, line, message)
	}
	fPath := strings.Split(fn.Name(), "/")
	funcName := fPath[len(fPath)-1]
	return fmt.Sprintf("[%s:%d - %s] %s", filename, line, funcName, message)
}
