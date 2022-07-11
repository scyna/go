package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func CreateModule(LOG scyna.Logger, module *model.Module) *scyna.Error {
	if applied, err := qb.Insert("scyna.module").
		Columns("org_code", "code", "description", "secret").
		Unique().Query(scyna.DB).
		BindStruct(module).
		ExecCASRelease(); !applied {
		if err == nil {
			return model.ORGANIZATION_EXISTED
		} else {
			LOG.Info("Can not create organization " + module.Code + " : " + err.Error())
			return scyna.SERVER_ERROR
		}
	}
	return scyna.OK
}
