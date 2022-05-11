package trace

import "github.com/scylladb/gocqlx/v2"

var qTraceInsert *gocqlx.Queryx
var qTraceInsertBatch *gocqlx.Queryx
var qTagInsert *gocqlx.Queryx
var qServiceTag *gocqlx.Queryx

func Init() {
	/*TODO*/
}

func Release() {
	/*TODO*/
}
