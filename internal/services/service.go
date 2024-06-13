package service

import (
	"github.com/AlmazHb/Auth/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

// type Claims struct {
// 	UserId string `json:"user-id"`
// 	jwt.StandardClaims
// }
