package repository

import (
	"github.com/gocql/gocql"
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/scyna"
)

const KEY_SPACE = "ex_Account"
const TABLE_NAME = "account"

func CreateAccount(a model.Account) *scyna.Error {
	batch := scyna.DB.NewBatch(gocql.LoggedBatch)
	batch.Query("INSERT INTO "+KEY_SPACE+"."+TABLE_NAME+"(entity_id, email,tel,first_name,last_name,name,language,created,updated,active,metadata)"+
		"VALUES(?,?,?,?,?,?,?,?,?,?,?)", a.ID, a.Email, a.Tel, a.FirstName, a.LastName, a.Name, a.Language, a.Created, a.Updated, true, a.Metadata)

	if err := scyna.DB.ExecuteBatch(batch); err != nil {
		LOG.Error("Can not execute batch " + err.Error())
		return model.ACCOUNT_EXISTS
	}
	return nil
}
