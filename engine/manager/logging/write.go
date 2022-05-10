package logging

import (
	"log"

	"github.com/scyna/go/scyna"
)

func Write(signal *scyna.LogCreatedSignal) {
	log.Print(signal.Text)
	scyna.AddLog(scyna.LogData{
		ID:       signal.ID,
		Sequence: signal.SEQ,
		Level:    scyna.LogLevel(signal.Level),
		Message:  signal.Text,
		Session:  signal.Session,
	})
}
