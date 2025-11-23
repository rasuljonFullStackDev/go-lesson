package main

import (
	"fmt"
	"net/http"

	"25.architect/internal/handlers"
	"25.architect/pkg/helpers"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	helpers.Success(w, "Hello World", http.StatusOK)
}
func test(w http.ResponseWriter, r *http.Request) {
	helpers.Success(w, "test", http.StatusOK)
}
func main() {
	// config.ConnectDB()
	// config.DB.AutoMigrate(&models.Product{})

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/test", test).Methods("GET")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	fmt.Println("🚀 Server http://localhost:8080 da ishlayapti...")
	http.ListenAndServe(":8080", r)
}
