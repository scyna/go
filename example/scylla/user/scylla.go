package user

import (
	"log"
	"sync"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type ScyllaRepository struct {
	GetQueries *scyna.QueryPool
}

var Repository IRepository

func InitScyllaRepository() {
	Repository = &ScyllaRepository{
		GetQueries: &scyna.QueryPool{sync.Pool{New: func() interface{} { return newGetQuery() }}},
	}
}

func (r *ScyllaRepository) Create(LOG scyna.Logger, user *User) *scyna.Error {
	if err := qb.Insert("ex.user").
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		BindStruct(user).
		ExecRelease(); err == nil {
		return nil
	} else {
		log.Print(err)
	}
	return scyna.SERVER_ERROR
}

func (r *ScyllaRepository) Exist(LOG scyna.Logger, email string) *scyna.Error {
	var id uint64
	if err := qb.Select("ex.user").
		Columns("id").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).
		Bind(email).
		GetRelease(&id); err == nil {
		return nil
	}
	return USER_NOT_EXISTED
}

func newGetQuery() *gocqlx.Queryx {
	return qb.Select("ex.user").
		Columns("id", "name", "email", "password").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB)
}

func (r *ScyllaRepository) GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User) {
	var query = r.GetQueries.GetQuery()
	defer r.GetQueries.PutQuery(query)

	var user User
	if err := query.Bind(email).GetRelease(&user); err == nil {
		return nil, &user
	} else {
		log.Print(err)
	}
	return USER_NOT_EXISTED, nil
}
