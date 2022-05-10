package generator

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

const snPartitionSize = 500

func GetSN(s *scyna.Service, request *scyna.GetSNRequest) {
	log.Print("Receive GetSNRequest")

	for i := 0; i < tryCount; i++ {
		if bucket := nextBucket(request.Key); bucket != nil {
			s.Done(bucket)
			return
		}
	}

	s.Error(scyna.SERVER_ERROR)
}

func nextBucket(key string) *scyna.GetSNResponse {
	prefix := time.Now().Unix() / (60 * 60 * 24)
	seed := 0
	if err := qb.Select("scyna.gen_sn").
		Columns("seed").
		Where(qb.Eq("key"), qb.Eq("prefix")).
		Limit(1).
		Query(scyna.DB).
		Bind(key, prefix).
		GetRelease(&seed); err == nil {
		seed += snPartitionSize
	} else {
		log.Print("OneID:", err)
	}

	if applied, err := qb.Insert("scyna.gen_sn").
		Columns("key", "prefix", "seed").
		Unique().
		Query(scyna.DB).
		Bind(key, prefix, seed).
		ExecCASRelease(); applied {
		return &scyna.GetSNResponse{
			Prefix: uint32(prefix),
			Start:  uint64(seed) + 1,
			End:    uint64(seed) + snPartitionSize,
		}
	} else {
		if err != nil {
			log.Print("nextBucket: insert seed: ", err.Error())
		}
	}
	return nil
}
