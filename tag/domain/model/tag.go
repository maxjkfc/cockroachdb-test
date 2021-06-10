package model

import "time"

type Record struct {
	User      User      `json:"user"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
