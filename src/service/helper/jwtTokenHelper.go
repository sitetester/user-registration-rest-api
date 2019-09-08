package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Concepts taken from  https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/

const AppSecret = "123MyRan&^dom#$#Str*&ing!"

// Create the JWT key used to create the signature
var jwtKey = []byte(AppSecret)

type JWTTokenHelper struct{}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string
	jwt.StandardClaims
}

func (jwtTokenHelper JWTTokenHelper) GenerateToken(email string) string {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	fmt.Println(AppSecret)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Println("Unable to sign token with secret key.")
	}

	return tokenString
}

func (jwtTokenHelper JWTTokenHelper) CheckToken(tokenString string) bool {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the tokenString is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
		}
		return false
	}

	return token.Valid
}
