package manager

import "github.com/scyna/go/scyna"

const (
	ENGINE_CONTEXT = "scyna.engine"
	ENGINE_SECRET  = "123456"
)

var DefaultConfig *scyna.Configuration = &scyna.Configuration{
	NatsUrl:      "127.0.0.1",
	NatsUsername: "",
	NatsPassword: "",
	DBHost:       "127.0.0.1",
	DBUsername:   "",
	DBPassword:   "",
	DBLocation:   "",
}
