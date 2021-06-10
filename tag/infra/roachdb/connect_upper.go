package roachdb

import (
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

type upperDatabase struct {
	db db.Session
}

// NewWithUpper -
func NewWithUpper(host, database, user, password string) (db.Session, error) {

	settings := cockroachdb.ConnectionURL{
		Host:     host,
		Database: database,
		User:     user,
		Password: password,
		Options: map[string]string{
			"sslmode": "disable",
		},
	}

	session, err := cockroachdb.Open(settings)
	if err != nil {
		return nil, err
	}

	session.SetMaxOpenConns(10)
	session.SetConnMaxLifetime(60 * time.Second)
	session.SetMaxIdleConns(3)

	return session, nil
}
