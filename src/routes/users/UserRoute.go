package users

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"io"
	"io/ioutil"
	"net/http"
	"user-registration-rest-api/src/entity"
	"user-registration-rest-api/src/service/db"
	"user-registration-rest-api/src/service/helper"
	"user-registration-rest-api/src/service/response"
)

type RouteUsers struct{}

type JWTTokenData struct {
	Token string `json:"token"`
}

func (u RouteUsers) LoginHandler(w http.ResponseWriter, req *http.Request) {
	var apiResponse response.ApiResponse
	var user entity.User
	userData := make(map[string]string)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		panic(err)
	}

	if len(userData["email"]) == 0 || len(userData["password"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		apiResponse.Error(w, req, "Missing email/password.")
		return
	}

	queryCount := 0
	db.GetDb().Where("email = ?", userData["email"]).First(&user).Count(&queryCount)
	if queryCount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		apiResponse.Error(w, req, "There is no user registered with this email.")
		return
	}

	var bcryptHelper helper.BcryptHelper
	if !bcryptHelper.CheckPasswordHash(userData["password"], user.Password) {
		w.WriteHeader(http.StatusBadRequest)
		apiResponse.Error(w, req, "Invalid password.")
		return
	}

	var jWTTokenHelper helper.JWTTokenHelper
	tokenString := jWTTokenHelper.GenerateToken(user.Email)

	jwtTokenData := JWTTokenData{Token: tokenString}
	jwtTokenDataBytes, err := json.Marshal(jwtTokenData)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.WriteString(w, string(jwtTokenDataBytes))
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
		w.WriteHeader(http.StatusBadRequest)
		apiResponse.Error(w, req, "Invalid email.")
		return
	}

	if govalidator.IsNull(user.Password) {
		w.WriteHeader(http.StatusBadRequest)
		apiResponse.Error(w, req, "Password is required.")
		return
	}

	queryCount := 0
	db.Where("email = ?", user.Email).First(&user).Count(&queryCount)
	if queryCount > 0 {
		w.WriteHeader(http.StatusBadRequest)
		apiResponse.Error(w, req, "Email already taken.")
		return
	}

	var bcryptHelper helper.BcryptHelper
	user.Password = bcryptHelper.HashPassword(user.Password)
	db.Create(&user)
	w.WriteHeader(http.StatusCreated)
	apiResponse.Success(w, req, "User registered successfully.")
}

// TODO: check request token in HTTP header (e.g Authorization Bearer XXX) or pass as cookie
func (u RouteUsers) DashboardHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var jWTTokenData JWTTokenData
	err := decoder.Decode(&jWTTokenData)
	if err != nil {
		panic(err)
	}

	var jWTTokenHelper helper.JWTTokenHelper
	tokenValid := jWTTokenHelper.CheckToken(jWTTokenData.Token)

	if tokenValid {
		io.WriteString(w, "Welcome !")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Invalid token.")
	}
}

func (u RouteUsers) UploadImageHandler(w http.ResponseWriter, req *http.Request) {
	var apiResponse response.ApiResponse
	const MaxFileUploadSize = 5000000
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	req.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, fileHeader, err := req.FormFile("image")
	if err != nil {
		apiResponse.Error(w, req, "Error retrieving file with `image` key.")
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
	fmt.Printf("File Size: %+v\n", fileHeader.Size)
	fmt.Printf("MIME Header: %+v\n", fileHeader.Header)

	if fileHeader.Size > MaxFileUploadSize {
		apiResponse.Error(w, req, "Max allowed file upload size: 5MB")
		return
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("public", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
