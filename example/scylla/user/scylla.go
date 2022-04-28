package user

import (
	"sync"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type ScyllaRepository struct {
	qInsert scyna.QueryPool
	qSelect scyna.QueryPool
}

var Repository IRepository

func InitScyllaRepository() {
	Repository = &ScyllaRepository{
		qInsert: scyna.QueryPool{sync.Pool{New: func() interface{} { return newInsertQuery() }}},
		qSelect: scyna.QueryPool{sync.Pool{New: func() interface{} { return newSelectQuery() }}},
		/*TODO: others query pools here*/
	}
}

func newInsertQuery() *gocqlx.Queryx {
	return qb.Insert("ex.user").Columns("id", "name", "email", "password").Query(scyna.DB)
}

func newSelectQuery() *gocqlx.Queryx {
	return qb.Select("ex.user").Columns("id").Where(qb.Eq("email")).Limit(1).Query(scyna.DB)
}

func (r *ScyllaRepository) Create(LOG scyna.Logger, user *User) *scyna.Error {
	var query = r.qInsert.GetQuery()
	defer r.qInsert.PutQuery(query)
	if err := query.Bind(user).Exec(); err == nil {
		return nil
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
		SelectRelease(&id); err == nil {
		return nil
	}
	return USER_NOT_EXISTED
}

func (r *ScyllaRepository) GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User) {
	var user User
	if err := qb.Select("ex.user").
		Columns("id").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).
		Bind(email).
		SelectRelease(&user); err == nil {
		return nil, &user
	}
	return USER_NOT_EXISTED, nil
}

func (r *ScyllaRepository) Release() {
	for q := r.qInsert.GetQuery(); q != nil; {
		q.Release()
	}
	for q := r.qSelect.GetQuery(); q != nil; {
		q.Release()
	}
}
