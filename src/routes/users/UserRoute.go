package users

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
	"user-registration-rest-api/src/entity"
	"user-registration-rest-api/src/service/db"
	"user-registration-rest-api/src/service/helper"
	"user-registration-rest-api/src/service/response"
)

type RouteUsers struct {
}

func (u RouteUsers) LoginHandler(w http.ResponseWriter, req *http.Request) {
	var resp response.ApiResponse
	var user entity.User
	userData := make(map[string]string)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		panic(err)
	}

	queryCount := 0
	db.GetDb().Where("email = ?", userData["email"]).First(&user).Count(&queryCount)
	if queryCount == 0 {
		resp.Error(w, req, "There is no user registered with this email.")
		return
	}

	var bcryptHelper helper.BcryptHelper
	if !bcryptHelper.CheckPasswordHash(userData["password"], user.Password) {
		resp.Error(w, req, "Invalid password.")
		return
	}
}

func (u RouteUsers) RegisterHandler(w http.ResponseWriter, req *http.Request) {
	db := db.GetDb()
	var apiResponse response.ApiResponse

	decoder := json.NewDecoder(req.Body)
	var user entity.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	if !govalidator.IsEmail(user.Email) {
		apiResponse.Error(w, req, "Invalid email.")
		return
	}

	if govalidator.IsNull(user.Password) {
		apiResponse.Error(w, req, "Password is required.")
		return
	}

	queryCount := 0
	db.Where("email = ?", user.Email).First(&user).Count(&queryCount)
	if queryCount > 0 {
		apiResponse.Error(w, req, "Email already taken.")
		return
	}

	var bcryptHelper helper.BcryptHelper
	user.Password = bcryptHelper.HashPassword(user.Password)
	db.Create(&user)
	apiResponse.Success(w, req, "User registered successfully.")
	// return
}
