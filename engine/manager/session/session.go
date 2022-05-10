package session

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func Init(module string, secret string) {
	if id, err := newSession(module, secret); err != scyna.OK {
		log.Fatal("Error in create session")
	} else {
		scyna.Session = scyna.NewSession(id)
	}
}

func newSession(module string, secret string) (uint64, *scyna.Error) {
	log.Print("Creating session for module: ", module)
	var secret_ string
	if err := qb.Select("scyna.module").
		Columns("secret").
		Where(qb.Eq("code")).
		Limit(1).
		Query(scyna.DB).
		Bind(module).
		GetRelease(&secret_); err != nil {
		log.Print("Module not existed: ", err.Error())
		return 0, scyna.MODULE_NOT_EXISTED
	}

	if secret != secret_ {
		log.Print("Module secret is not correct")
		return 0, scyna.PERMISSION_ERROR
	}

	sid := scyna.ID.Next()
	now := time.Now()

	if err := qb.Insert("scyna.session").
		Columns("id", "module_code", "start", "last_update").
		Query(scyna.DB).
		Bind(sid, module, now, now).
		ExecRelease(); err != nil {
		log.Print("Can not save session to database:", err)
		return 0, scyna.SERVER_ERROR
	}

	log.Print("Session Created")
	return sid, scyna.OK
}
