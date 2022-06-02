package scyna

import (
	"sync"

	"github.com/scylladb/gocqlx/v2"
)

type QueryPool struct {
	sync.Pool
}

type QueryCreator func() *gocqlx.Queryx

func NewQueryPool(creator QueryCreator) *QueryPool {
	return &QueryPool{sync.Pool{New: func() interface{} { return creator() }}}
}

func (q *QueryPool) GetQuery() *gocqlx.Queryx {
	query, _ := q.Get().(*gocqlx.Queryx)
	return query
}

func (q *QueryPool) PutQuery(query *gocqlx.Queryx) {
	q.Put(query)
}
