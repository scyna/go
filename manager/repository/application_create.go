package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func CreateApplication(LOG scyna.Logger, app *model.Application) *scyna.Error {
	if applied, err := qb.Insert("scyna.application").
		Columns("org_code", "code", "auth", "name").
		Unique().Query(scyna.DB).
		BindStruct(app).
		ExecCASRelease(); !applied {
		if err == nil {
			return model.APPLICATION_EXISTED
		} else {
			LOG.Info("Can not create application " + app.Code + " : " + err.Error())
			return scyna.SERVER_ERROR
		}
	}
	return nil
}
