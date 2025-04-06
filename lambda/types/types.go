package types

import (
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(ru RegisterUser) (User, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(ru.Password), 10)
	if err != nil {
		return User{}, err
	}
	return User{
		Username:     ru.Username,
		PasswordHash: string(hp),
	}, nil
}

func ValidateUser(hp, pp string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(pp))
	return err == nil
}
