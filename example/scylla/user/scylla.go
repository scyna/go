package user

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type ScyllaRepository struct {
}

var Repository IRepository

func InitScyllaRepository() {
	Repository = &ScyllaRepository{}
}

func (r *ScyllaRepository) Create(LOG scyna.Logger, user *User) *scyna.Error {
	if err := qb.Insert("ex.user").
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		Bind(user).
		ExecRelease(); err == nil {
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
	return nil, &User{}
}
