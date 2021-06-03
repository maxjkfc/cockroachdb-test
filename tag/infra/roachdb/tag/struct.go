package tag

import (
	"time"

	"github.com/upper/db/v4/adapter/cockroachdb"
)

type Record struct {
	ID        string            `db:"id,omitempty"`
	UserInfo  cockroachdb.JSONB `db:"user_info"`
	Tags      []string          `db:"tags"`
	CreatedAt time.Time         `db:"createdat"`
}
