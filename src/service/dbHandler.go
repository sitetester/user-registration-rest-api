package service

import (
	"github.com/jinzhu/gorm"
	"user-registration-rest-api/src/entity"
)

func ManageDb() {
	db, err := gorm.Open("sqlite3", "./db/app.db")
	if err != nil {
		panic("Failed to connect database!")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&entity.User{})
}
