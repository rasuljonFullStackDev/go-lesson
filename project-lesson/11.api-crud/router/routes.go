package router

import (
	"github.com/gorilla/mux"
	"11.api-crud/handlers"
	"11.api-crud/middleware"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	// Auth
	api.HandleFunc("/register", handlers.Register).Methods("POST")
	api.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	protected.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	protected.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	return r
}
