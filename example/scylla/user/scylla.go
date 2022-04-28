package user

import (
	"log"

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

func (r *ScyllaRepository) GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User) {
	var user User
	if err := qb.Select("ex.user").
		Columns("id", "name", "email", "password").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).
		Bind(email).
		GetRelease(&user); err == nil {
		return nil, &user
	} else {
		log.Print(err)
	}
	return USER_NOT_EXISTED, nil
}
