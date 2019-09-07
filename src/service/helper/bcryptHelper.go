package helper

import "golang.org/x/crypto/bcrypt"

type BcryptHelper struct {
}

func (bcryptHelper BcryptHelper) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func (bcryptHelper BcryptHelper) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
