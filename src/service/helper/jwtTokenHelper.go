package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const AppSecret = "123MyRan&^dom#$#Str*&ing!"

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
	var jwtKey = []byte(AppSecret)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Println("Unable to sign token with secret key.")
	}

	return tokenString
}
