package scyna_data

import (
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func WriteSetting(module string, key string, value string) {
	if applied, err := qb.Insert("scyna.setting").
		Columns("module_code", "key", "value").
		Unique().
		Query(scyna.DB).
		Bind(module, key, value).
		ExecCASRelease(); !applied {
		if err != nil {
			log.Fatal("Error in write setting:", err.Error())
			return
		}
	}

	scyna.SendCommand(scyna.SETTING_UPDATE_CHANNEL+module, &scyna.SettingUpdatedSignal{
		Key:   key,
		Value: value,
	})
}
