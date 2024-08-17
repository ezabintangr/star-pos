package encrypts

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type HashInterface interface {
	CheckPasswordHash(hashed string, password string) bool
	HashPassword(password string) (string, error)
}

type hash struct{}

func NewHashService() HashInterface {
	return &hash{}
}

// CheckPasswordHash implements HashInterface.
func (h *hash) CheckPasswordHash(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

// HashPassword implements HashInterface.
func (h *hash) HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error hash password: " + err.Error())
		return "", err
	}
	return string(result), nil
}
