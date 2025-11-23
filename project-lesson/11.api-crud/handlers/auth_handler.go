package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"11.api-crud/config"
	"11.api-crud/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("supersecret") // .env orqali olish mumkin

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email va parol kerak", http.StatusBadRequest)
		return
	}

	// Parolni hash qilish
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Hash xatosi", http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Foydalanuvchi bazada mavjud", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Foydalanuvchi ro‘yxatdan o‘tdi",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "Foydalanuvchi topilmadi", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Parol xato", http.StatusUnauthorized)
		return
	}

	// JWT yaratish
	expiration := time.Now().Add(1 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     expiration.Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Token yaratishda xato", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
