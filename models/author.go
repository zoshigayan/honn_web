package models

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Books []Book `gorm:"many2many:book_authors;"`
}
