package db

import (
	"github.com/jinzhu/gorm"
	"user-registration-rest-api/src/entity"
)

type Handler struct{}

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./db/app.db")
	if err != nil {
		panic("Failed to connect database!")
	}
	db.LogMode(true)

	// https://stackoverflow.com/questions/27241187/golang-sql-database-open-and-close
	// no need to close db connection
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&entity.User{})

	return db
}
