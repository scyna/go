package user

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type repository struct {
	GetQueries *scyna.QueryPool /*for high frequence queries*/
}

var Repository *repository = &repository{GetQueries: scyna.NewQueryPool(newGetQuery)}

func (r *repository) PrepareCreate(cmd *scyna.Command, user *User) {
	cmd.Batch.Query("INSERT INTO ex.user(id, name, email, password) VALUES(?,?,?,?)",
		user.ID,
		user.Name,
		user.Email,
		user.Password)
}

func (r *repository) Create(LOG scyna.Logger, user *User) *scyna.Error {
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

func (r *repository) GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User) {
	var user User
	if err := qb.Select("ex.user").
		Columns("id", "name", "email", "password").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).Bind(email).Get(&user); err != nil {
		LOG.Error(err.Error())
		return USER_NOT_EXISTED, nil
	}
	return nil, &user
}

func (r *repository) ListFriend(LOG scyna.Logger, uid uint64) (*scyna.Error, []*User) {
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
