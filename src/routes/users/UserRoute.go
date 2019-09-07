package users

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"io"
	"net/http"
	"user-registration-rest-api/src/entity"
	"user-registration-rest-api/src/service/helper"
	"user-registration-rest-api/src/service/response"
)

type RouteUsers struct {
}

func (u RouteUsers) LoginHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user entity.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	io.WriteString(w, user.Email+" ==== "+user.Password)
}

func (u RouteUsers) RegisterHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user entity.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	if !govalidator.IsEmail(user.Email) {
		var resp response.ErrorResponse
		resp.Error(w, req, "Invalid email!")
		return
	}

	if govalidator.IsNull(user.Password) {
		var resp response.ErrorResponse
		resp.Error(w, req, "Password is required")
		return
	}

	var bcryptHelper helper.BcryptHelper
	io.WriteString(w, bcryptHelper.HashPassword(user.Password))
}
