package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Thumbnail   string
	Title       string
	Subtitle    string
	Authors     []Author `gorm:"many2many:book_authors;"`
	PublishedAt time.Time
	Description string
}
