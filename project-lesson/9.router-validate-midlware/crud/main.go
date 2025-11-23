package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET – barcha mahsulotlar")
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST – yangi mahsulot yaratildi")
}

func update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "PUT – mahsulot yangilandi: %s", id)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "DELETE – mahsulot o‘chirildi: %s", id)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", getAll).Methods("GET")
	r.HandleFunc("/products", create).Methods("POST")
	r.HandleFunc("/products/{id}", update).Methods("PUT")
	r.HandleFunc("/products/{id}", delete).Methods("DELETE")

	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
