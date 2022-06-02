package scyna

import (
	"sync"

	"github.com/scylladb/gocqlx/v2"
)

type QueryPool struct {
	sync.Pool
}

type NewQuery func() *gocqlx.Queryx

func NewQueryPool(newQuery NewQuery) *QueryPool {
	return &QueryPool{sync.Pool{New: func() interface{} { return newQuery() }}}
}

func (q *QueryPool) GetQuery() *gocqlx.Queryx {
	query, _ := q.Get().(*gocqlx.Queryx)
	return query
}

func (q *QueryPool) PutQuery(query *gocqlx.Queryx) {
	q.Put(query)
}
