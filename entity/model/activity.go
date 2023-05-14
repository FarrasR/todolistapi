package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ActivityID uint           `gorm:"primaryKey" json:"id"`
	Email      string         `gorm:"type:varchar(200);not null" json:"email"`
	Title      string         `gorm:"type:varchar(100);not null" json:"title"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
