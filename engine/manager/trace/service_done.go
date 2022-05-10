package trace

import (
	"log"

	"github.com/scyna/go/scyna"
)

func ServiceDone(signal *scyna.TagCreatedSignal) {
	log.Print("Service Done")
	/* TODO */
}
