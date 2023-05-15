package database

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V2023051500001 = gormigrate.Migration{
	ID: "V2023051500001",
	Migrate: func(tx *gorm.DB) error {
		type Activity struct {
			ActivityID uint           `gorm:"primaryKey" json:"id"`
			Email      string         `gorm:"type:varchar(200);not null" json:"email"`
			Title      string         `gorm:"type:varchar(100);not null" json:"title"`
			CreatedAt  time.Time      `json:"created_at"`
			UpdatedAt  time.Time      `json:"updated_at"`
			DeletedAt  gorm.DeletedAt `json:"deleted_at"`
		}
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

		return tx.AutoMigrate(&Activity{}, &Todo{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("activities", "todos")
	},
}
