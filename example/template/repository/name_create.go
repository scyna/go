package repository

import (
	"template/model"

	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

func CreateName(LOG scyna.Logger, name string) *scyna.Error {
	batch := scyna.DB.NewBatch(gocql.LoggedBatch)
	batch.Query("INSERT INTO template.name(name) VALUES(?);", name)
	if err := scyna.DB.ExecuteBatch(batch); err != nil {
		LOG.Error(err.Error())
		return model.NAME_EXISTED
	}

	return nil
}
