package dto

import "time"

type ShortenResponse struct {
	ShortUrl  string    `json:"short_url"`
	LongUrl   string    `json:"long_url"`
	ExpiresAt time.Time `json:"expires_at"`
}
