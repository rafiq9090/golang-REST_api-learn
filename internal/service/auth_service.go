package service

import (
	"errors"
	"task-api/internal/config"
	"task-api/internal/models"
	"task-api/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(config.App.JWTSecret)

type AuthService struct{}

var Auth = AuthService{}

func (AuthService) Register(user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return repository.Auth.Register(user)
}

func (AuthService) Login(email string, password string) (*models.User, string, error) {
	user, err := repository.Auth.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("Invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("Invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)

	return user, tokenString, nil

}
