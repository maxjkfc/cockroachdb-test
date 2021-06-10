package roachdb

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewWithPgx(host, database, user, password string) (*pgxpool.Pool, error) {

	url := "postgres://"

	if user != "" {
		url += user
	}

	if password != "" {
		url += ":" + password
	}

	if host != "" {
		if user != "" {
			url += "@"
		}
		url += host
	}

	url += "?sslmode=disable"

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic(err)
	}

	config.MaxConnIdleTime = time.Second * 60
	config.MaxConns = 10
	config.MinConns = 3

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("connect failed")
	}

	return conn, err
}
