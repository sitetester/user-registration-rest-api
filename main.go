package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"user-registration-rest-api/src/service"
)

const AppSecret = "123MyRan&^dom#$#Str*&ing!"

func main() {
	service.ManageRoutes()
}
