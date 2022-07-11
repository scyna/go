package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func CreateClient(LOG scyna.Logger, client *model.Client) *scyna.Error {
	if applied, err := qb.Insert("scyna.client").
		Columns("org_code", "id", "secret").
		Unique().Query(scyna.DB).
		BindStruct(client).
		ExecCASRelease(); !applied {
		if err == nil {
			return model.CLIENT_EXISTED
		} else {
			LOG.Info("Can not create client " + client.ID + " : " + err.Error())
			return scyna.SERVER_ERROR
		}
	}
	return nil
}
