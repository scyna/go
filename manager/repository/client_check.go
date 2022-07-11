package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
	"log"
)

func CheckClient(org string, id string) bool {
	if err := qb.Select("scyna.client").
		Columns("id").
		Where(qb.Eq("org_code"), qb.Eq("id")).
		Query(scyna.DB).
		Bind(org, id).
		GetRelease(&id); err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
