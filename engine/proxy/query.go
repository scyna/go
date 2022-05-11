package proxy

import (
	"github.com/scyna/go/scyna"
	"sync"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type Query struct {
	Authenticate *gocqlx.Queryx
}

type QueryPool struct {
	sync.Pool
}

func NewQuery() *Query {
	return &Query{
		Authenticate: qb.Select("scyna.client_use_service").
			Columns("service_url").
			Where(qb.Eq("client_id"), qb.Eq("service_url")).
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
