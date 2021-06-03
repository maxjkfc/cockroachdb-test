package model

import "time"

type Filter struct {
	User    User
	Tags    []string
	StartAt time.Time
	EndAt   time.Time
	Page    int64
	Size    int64
}
