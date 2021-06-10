package tag

import (
	"context"
	"tag/domain/model"
	"tag/usecase/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjarratt/babble"
)

var _conn_upper repository.TagRepository
var _conn_pgx repository.TagRepository

func Test_New(t *testing.T) {
	t.Run("new connect with upper", func(t *testing.T) {
		conn, err := NewRepositoryWithUpper("localhost:26257", "tag", "admin", "")
		assert.NoError(t, err)
		assert.NotEmpty(t, conn)

		_conn_upper = conn
	})
	t.Run("new connect with pxgpool", func(t *testing.T) {
		conn, err := NewRepositoryWithPgx("localhost:26257", "tag", "admin", "")
		assert.NoError(t, err)
		assert.NotEmpty(t, conn)

		_conn_pgx = conn
	})
}

func Test_Create(t *testing.T) {
	t.Run("create records with upper", func(t *testing.T) {
		b := babble.NewBabbler()
		b.Separator = ""
		b.Count = 1

		r := model.Record{
			User: model.User{
				Name: "upper",
			},
			Tag: b.Babble(),
		}

		err := _conn_upper.Create(context.Background(), r)
		assert.NoError(t, err)
	})
	t.Run("create records with pgx", func(t *testing.T) {
		b := babble.NewBabbler()
		b.Separator = ""
		b.Count = 1

		r := model.Record{
			User: model.User{
				Name: "pgx",
			},
			Tag: b.Babble(),
		}

		err := _conn_pgx.Create(context.Background(), r)
		assert.NoError(t, err)
	})
}

func Test_Search(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		filter := model.Filter{
			User: model.User{
				Name: "upper",
			},
		}
		total, ans := _conn_upper.Search(context.Background(), filter)
		assert.NotZero(t, total)
		assert.NotEmpty(t, ans)
	})
}
