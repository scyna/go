package generator

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

const idPartitionSize = 1000
const tryCount = 10

func Init() {
	for i := 0; i < tryCount; i++ {
		if ok, prefix, start, end := allocate(); ok {
			scyna.ID.Reset(prefix, end, start)
			return
		}
	}
	log.Fatal("Can not init id generator")
}

func GetID(s *scyna.Service, request *scyna.GetIDRequest) {
	log.Print("Receive GetIDRequest")
	for i := 0; i < tryCount; i++ {
		if ok, prefix, start, end := allocate(); ok {
			s.Done(&scyna.GetIDResponse{
				Prefix: prefix,
				Start:  start,
				End:    end,
			})
			log.Print("Return GetIDResponse")
			return
		}
	}
	s.Error(scyna.SERVER_ERROR)
}

func allocate() (ok bool, prefix uint32, start uint64, end uint64) {
	p := time.Now().Unix() / (60 * 60 * 24)
	ok = false

	seed := 0
	if err := qb.Select("scyna.gen_id").
		Columns("seed").
		Where(qb.Eq("prefix")).
		Limit(1).
		Query(scyna.DB).
		Bind(p).
		GetRelease(&seed); err == nil {
		seed += idPartitionSize
	} else {
		log.Println("generator.allocate: get seed: " + err.Error())
	}

	if applied, err := qb.Insert("scyna.gen_id").
		Columns("prefix", "seed").
		Unique().
		Query(scyna.DB).
		Bind(p, seed).
		ExecCASRelease(); applied {
		prefix = uint32(p)
		start = uint64(seed) + 1
		end = uint64(seed) + idPartitionSize
		ok = true
		return
	} else {
		if err != nil {
			log.Println("generator.allocate: insert: " + err.Error())
		} else {
			log.Println("generator.allocate: cannot insert ")
		}
	}
	return
}
