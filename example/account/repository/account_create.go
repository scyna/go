package repository

import (
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/scyna"
)

const KEY_SPACE = "ex_Account"
const TABLE_NAME = "account"

func CreateAccount(LOG scyna.Logger, a model.Account) *scyna.Error {
	//batch := scyna.DB.NewBatch(gocql.LoggedBatch)
	//batch.Query("INSERT INTO "+KEY_SPACE+"."+TABLE_NAME+"(entity_id,email,name,password)"+
	//	"VALUES(?,?,?,?)", a.ID, a.Email, a.Name, a.Password)
	//
	//if err := scyna.DB.ExecuteBatch(batch); err != nil {
	//	LOG.Error("Can not execute batch " + err.Error())
	//	return model.ACCOUNT_EXISTS
	//}
	return nil
}
