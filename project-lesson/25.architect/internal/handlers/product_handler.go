package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"25.architect/internal/models"
	"25.architect/internal/services"

	"github.com/gorilla/mux"
)

var productService = services.ProductService{}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	p.Name = r.FormValue("name")
	p.Price = 1500 // test uchun

	if err := productService.Create(&p, r); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := productService.FindAll(&products); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var p models.Product
	if err := productService.Delete(uint(id), &p); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
