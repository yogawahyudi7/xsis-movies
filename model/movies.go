package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Id          uint    `gorm:"primaryKey;AUTO_INCREMENT;column:id"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Rating      float64 `gorm:"column:rating"`
	Image       string  `gorm:"column:image"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
