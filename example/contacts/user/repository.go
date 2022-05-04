package user

import (
	"sync"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type ScyllaRepository struct {
	GetQueries *scyna.QueryPool /*for high frequence queries*/
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
		ExecRelease(); err != nil {
		LOG.Error(err.Error())
		return scyna.SERVER_ERROR
	}
	return nil
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
	if err := query.Bind(email).Get(&user); err != nil {
		LOG.Error(err.Error())
		return USER_NOT_EXISTED, nil
	}
	return nil, &user
}

func (r *ScyllaRepository) ListFriend(LOG scyna.Logger, uid uint64) (*scyna.Error, []*User) {
	var friends []uint64
	var ret []*User

	if err := qb.Select("ex.has_friend").
		Columns("friend_id").
		Where(qb.Eq("user_id")).
		Limit(1).
		Query(scyna.DB).Bind(uid).SelectRelease(friends); err != nil {
		return scyna.SERVER_ERROR, ret
	}

	if len(friends) == 0 {
		return nil, ret
	}

	ret = make([]*User, len(friends))

	qSelect := qb.Select("ex.user").
		Columns("id", "name", "email").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB)

	for i, id := range friends {
		qSelect.Bind(id).Get(ret[i])
	}

	qSelect.Release()
	return nil, ret
}
