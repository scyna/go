package session

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

//https://tldp.org/LDP/abs/html/exitcodes.html

func End(signal *scyna.EndSessionSignal) {
	if applied, err := qb.Update("scyna.session").
		Set("end", "exit_code").
		Where(qb.Eq("id"), qb.Eq("module_code")).Existing().
		Query(scyna.DB).
		Bind(time.Now(), signal.Code, signal.ID, signal.Context).
		ExecCASRelease(); !applied {
		if err != nil {
			log.Print("Can not update EndSessionSignal:", err.Error())
		}
	}
}
