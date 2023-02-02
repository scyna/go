package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/scyna"
)

func GetAccountByEmail(LOG scyna.Logger, email string) (*model.Account, *scyna.Error) {

	var account *model.Account

	if err := qb.Select(KEY_SPACE+"."+TABLE_NAME).
		Columns("entity_id", "email", "name", "password").
		Where(qb.Eq("email")).
		Query(scyna.DB).
		Bind(email).
		GetRelease(&account); err != nil {
		LOG.Error(err.Error())
		return account, nil
	}
	return nil, model.ACCOUNT_NOT_FOUND
}
