package user

import (
	"encoding/json"
	"net/http"
	"project-structure/app/request/user"
	"project-structure/app/services/api/user"

	"gorm.io/gorm"
)

type UserController struct {
	Service user.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Service: user.UserService{DB: db},
	}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := c.Service.Create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
