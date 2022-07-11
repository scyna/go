package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func CreateOrganization(LOG scyna.Logger, org *model.Organization) *scyna.Error {
	if applied, err := qb.Insert("scyna.organization").
		Columns("code", "name", "password").
		Unique().Query(scyna.DB).
		Bind(org).
		ExecCASRelease(); !applied {
		if err == nil {
			return model.ORGANIZATION_EXISTED
		} else {
			LOG.Info("Can not create organization " + org.Code + " : " + err.Error())
			return scyna.SERVER_ERROR
		}
	}
	return scyna.OK
}
