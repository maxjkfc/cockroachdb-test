package roachdb

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upper/db/v4"
)

var (
	_upper          db.Session
	_collection     = "accounts"
	_upper_accounts = make(map[int]*Account, 0)
)

type Account struct {
	ID      uint64 `db:"id,omitempty"`
	Name    string `db:"name"`
	Balance int64  `db:"balance"`
}

func init() {

	for i := 0; i <= 10; i++ {
		_upper_accounts[i] = &Account{
			Name:    fmt.Sprintf("upper-account-%v", i),
			Balance: rand.Int63n(999999),
		}
	}
}

func TestConnect(t *testing.T) {

	t.Run("使用 upper 連結資料庫", func(t *testing.T) {
		session, err := NewWithUpper("localhost:26257", "bank", "maxroach", "")
		assert.NoError(t, err)
		assert.NotEmpty(t, session)
		_upper = session
	})

}

func TestCreateTable(t *testing.T) {
	t.Run("使用 upper 建立資料表", func(t *testing.T) {
		_, err := _upper.SQL().Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
		  ID SERIAL PRIMARY KEY,
		  Name VARCHAR(255) NOT NULL,
		  balance INT
		)
	  `)

		assert.NoError(t, err)
	})

}

func TestInsertData(t *testing.T) {
	t.Run("使用 upper 建立一筆資料", func(t *testing.T) {
		for _, v := range _upper_accounts {
			err := _upper.Collection(_collection).InsertReturning(v)
			assert.NoError(t, err)
		}

		accounts := make([]Account, 0)
		err := _upper.Collection(_collection).Find().All(&accounts)
		assert.NoError(t, err)

		for _, v := range accounts {
			fmt.Printf("\t accounts[%d]:  %s %d \n", v.ID, v.Name, v.Balance)
		}
	})
}
