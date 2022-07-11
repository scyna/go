package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func AddService(LOG scyna.Logger, client string, service string) *scyna.Error {
	if applied, err := qb.Insert("scyna.client_use_service").
		Columns("client_id", "service_url").
		Unique().Query(scyna.DB).
		Bind(client, service).
		ExecCASRelease(); !applied {
		if err == nil {
			return model.CLIENT_EXISTED
		} else {
			LOG.Info("Can not create client use service " + client + " : " + err.Error())
			return scyna.SERVER_ERROR
		}
	}
	return nil
}
