package dto

type ShortenRequest struct {
	LongURL string `json:"long_url" binding:"required,url"`
}
