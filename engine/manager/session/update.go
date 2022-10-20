package session

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func Update(signal *scyna.UpdateSessionSignal) {
	if applied, err := qb.Update("scyna.session").
		Set("last_update").
		Where(qb.Eq("id"), qb.Eq("context")).
		Existing().
		Query(scyna.DB).
		Bind(time.Now(), signal.ID, signal.Context).
		ExecCASRelease(); !applied {
		if err != nil {
			log.Print("Can not update UpdateSessionSignal: ", err.Error())
		}
	}
}
