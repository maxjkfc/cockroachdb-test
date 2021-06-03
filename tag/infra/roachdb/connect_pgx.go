package roachdb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type pgxDatabase struct {
	con *pgx.Conn
}

func NewWithPgx(host, database, user, password string) {

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

	config, err := pgx.ParseConfig(url)
	if err != nil {
		panic(err)
	}

	_ = config

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("connect failed")
	}

	_ = conn
}
