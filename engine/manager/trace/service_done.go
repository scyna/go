package trace

import (
	"log"

	"github.com/scyna/go/scyna"
)

func ServiceDone(signal *scyna.ServiceDoneSignal) {
	log.Print("Service Done")
	log.Print(signal.Request)
	log.Print(signal.Response)
	/* TODO */
}
