// Package tag provides ...
package tag

import (
	"tag/domain/model"
	"time"

	"github.com/upper/db/v4/adapter/cockroachdb"
)

func ConvertToDB(record ...model.Record) []Record {
	t := time.Now().Local()
	ans := make([]Record, len(record))
	for i, v := range record {
		ans[i] = Record{
			UserInfo:  cockroachdb.JSONB{V: v.User},
			Tag:       v.Tag,
			CreatedAt: t,
		}
	}
	return ans
}
