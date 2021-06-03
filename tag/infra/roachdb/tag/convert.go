// Package tag provides ...
package tag

import (
	"tag/domain/model"
	"time"

	"github.com/upper/db/v4/adapter/cockroachdb"
)

func ConvertToDB(record model.Record) Record {
	return Record{
		UserInfo:  cockroachdb.JSONB{V: record.User},
		Tags:      record.Tags,
		CreatedAt: time.Now().Local(),
	}
}
