package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	TodoID          uint   `gorm:"primaryKey"`
	ActivityGroupID uint   `gorm:"not null"`
	Title           string `gorm:"type:varchar(255);not null"`
	Priority        string
	IsActive        bool
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	UpdatedAt       time.Time
}
