package models

import "time"

type Url struct {
	Id        string
	Domain    string
	Code      string
	LongUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}
