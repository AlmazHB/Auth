package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlmazHb/Auth/internal/app/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth/tokens", h.GenerateTokens).Methods("POST")
	router.HandleFunc("/auth/refresh_tokens", h.RefreshTokens).Methods("POST")
	return router
}

func (h *Handler) GenerateTokens(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userID := r.Form.Get("GUID")

	// Генерируем и выдаем токены пользователю
	accessToken, refreshToken, err := h.s.GenerateToken(userID)
	if err != nil {
		http.Error(w, "Error of generate tokens", http.StatusInternalServerError)
		return
	}

	// Отправляем токены в ответе
	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.Fatal(err)
	}

}

func (h *Handler) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	var userId = "ushbdhd56"
	// Получаем refresh токен из параметров запроса
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при обработке запроса", http.StatusBadRequest)
		return
	}
	refreshToken := r.Form.Get("refresh_token")

	newAccessToken, newRefreshToken, err := h.s.RefreshToken(userId, refreshToken)
	if err != nil {
		http.Error(w, "Ошибка при обновлении токенов", http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}
	err = json.NewEncoder(w).Encode(response)
	log.Fatal(err)
}
