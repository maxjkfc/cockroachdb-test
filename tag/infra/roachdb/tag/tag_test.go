package tag

import (
	"context"
	"tag/domain/model"
	"tag/usecase/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _conn repository.TagRepository

func Test_New(t *testing.T) {
	t.Run("new connect", func(t *testing.T) {
		conn, err := NewRepositoryWithUpper("localhost:26257", "tag", "admin", "")
		assert.NoError(t, err)
		assert.NotEmpty(t, conn)

		_conn = conn
	})
}

func Test_Create(t *testing.T) {
	t.Run("create records", func(t *testing.T) {
		r := model.Record{
			User: model.User{
				Name: "max",
			},
			Tags: []string{
				"good", "like", "gooody",
			},
		}

		err := _conn.Create(context.Background(), r)
		assert.NoError(t, err)
	})
}
