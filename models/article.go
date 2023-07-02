package models

import (
	"time"
)

type Article struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
