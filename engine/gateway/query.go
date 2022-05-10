package gateway

import (
	"sync"

	"github.com/scyna/go/scyna"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type Query struct {
	GetSession   *gocqlx.Queryx
	CheckService *gocqlx.Queryx
}

type QueryPool struct {
	sync.Pool
}

func NewQuery() *Query {
	return &Query{
		CheckService: qb.Select("scyna.app_has_service").
			Columns("service_url").
			Where(qb.Eq("app_code"), qb.Eq("service_url")).
			Limit(1).
			Query(scyna.DB),

		GetSession: qb.Select("scyna.user_session").
			Columns("id", "expired", "app_code").
			Where(qb.Eq("code")).
			Limit(1).
			Query(scyna.DB),
	}
}

func (q *QueryPool) GetQuery() *Query {
	query, _ := q.Get().(*Query)
	return query
}

func (q *QueryPool) PutQuery(query *Query) {
	q.Put(query)
}

func NewQueryPool() QueryPool {
	return QueryPool{
		sync.Pool{
			New: func() interface{} { return NewQuery() },
		}}
}
