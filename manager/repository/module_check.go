package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func CheckModule(code string) bool {
	if err := qb.Select("scyna.module").
		Columns("code").
		Where(qb.Eq("code")).
		Query(scyna.DB).
		Bind(code).
		GetRelease(&code); err != nil {
		return false
	}
	return true
}
