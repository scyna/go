package scyna

import "github.com/nats-io/nats.go"

const es_TRY_COUNT = 10

type eventInStore struct {
	ID      uint64
	Subject string
	Data    []byte
}

func storeEvent(m *nats.Msg) bool {

	for i := 0; i < es_TRY_COUNT; i++ {
		if err, lastEvent := getLastEvent(); err == nil {
			if isDuplicate(lastEvent, m) {
				return true
			}

			if saveEvent(lastEvent.ID+1, m.Subject, m.Data) == nil {
				return true
			}
		}
	}

	return false
}

func getLastEvent() (*Error, *eventInStore) {
	/*TODO*/
	return nil, nil
}

func isDuplicate(e *eventInStore, m *nats.Msg) bool {
	/*TODO*/
	return false
}

func saveEvent(ID uint64, subject string, data []byte) *Error {
	/*TODO*/
	return nil
}
