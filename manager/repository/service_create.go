package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func CreateService(LOG scyna.Logger, service *model.Service) *scyna.Error {
	if applied, err := qb.Insert("scyna.service").
		Columns("module_code", "url", "description").
		Unique().Query(scyna.DB).
		BindStruct(service).
		ExecCASRelease(); !applied {
		if err == nil {
			return model.SERVICE_EXISTED
		} else {
			LOG.Info("Can not create service " + service.URL + " : " + err.Error())
			return scyna.SERVER_ERROR
		}
	}
	return nil
}
