package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"user-registration-rest-api/src/service"
)

func main() {
	service.ManageDb()
	service.ManageRoutes()
}
