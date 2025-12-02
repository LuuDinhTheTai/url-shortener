package model

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"`
	Domain    string    `gorm:"not null"`
	Code      string    `gorm:"not null;unique"`
	LongUrl   string    `gorm:"not null;type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	ExpiresAt time.Time `gorm:"not null"`
}
