package routes

import (
	"fmt"
	"net/http"
	userController "project-structure/app/controller/api/user"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) {
	r := mux.NewRouter()

	userCtrl := userController.NewUserController(db)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", userCtrl.Create).Methods(http.MethodPost)
	api.HandleFunc("/users", userCtrl.List).Methods(http.MethodGet)

	fmt.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", r)
}
