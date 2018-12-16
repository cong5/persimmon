package service

import (
	"github.com/cong5/persimmon/app/models"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (this *AuthService) Login(pwd string, user *models.Users) bool {
	pwd = strings.Trim(pwd, " ")
	password := strings.Trim(user.Password, " ")
	if user.Email == "" {
		return false
	}
	return this.CheckPasswordHash(pwd, password)
}

func (this *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (this *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
