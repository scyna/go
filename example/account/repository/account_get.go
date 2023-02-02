package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/scyna"
)

func GetAccountByID(LOG scyna.Logger, ID string) (*model.Account, *scyna.Error) {

	var account *model.Account

	if err := qb.Select(KEY_SPACE+"."+TABLE_NAME).
		Columns("entity_id", "email", "name", "password").
		Where(qb.Eq("entity_id")).
		Query(scyna.DB).
		Bind(ID).
		GetRelease(&account); err != nil {
		LOG.Error(err.Error())
		return account, nil
	}
	return nil, model.ACCOUNT_NOT_FOUND
}
