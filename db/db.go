package db

import (
	"github.com/zoshigayan/honn_web/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("honn.db"), &gorm.Config{})
	if err != nil {
		panic("DB接続に失敗しました。")
	}

	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
	log.Println("接続したで")
}

func DbManager() *gorm.DB {
	log.Println("とってくるで")
	return db
}
