package service

import (
	"time"

	"github.com/AlmazHb/Auth/internal/app/repository"
	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type Claims struct {
	UserId string `json:"user-id"`
	jwt.StandardClaims
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var jwtSecret = []byte("almaz.query")

func (s *Service) GenerateToken(userID string) (string, string, error) {

	// check user

	accessToken, err := generateAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (s *Service) RefreshToken(userId string, refreshtoken string) (string, string, error) {

	// refresh tokenin barlagy

	accessToken, err := generateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(userID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	})

	accessToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func generateRefreshToken(userID string) (string, error) {

	refreshToken := "random_refresh_token_for_user_" + userID

	// Хеширование refresh токена для безопасного хранения в базе данных
	// hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	// if err != nil {
	// 	return "", err
	// }

	return string(refreshToken), nil
}
