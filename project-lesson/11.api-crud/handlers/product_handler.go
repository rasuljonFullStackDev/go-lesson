package handlers

import (
	"encoding/json"
	"fmt"
	"11.api-crud/config"
	"11.api-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := config.DB.Model(&models.Product{}).Preload("Category")

	// filter parametrlari
	name := r.URL.Query().Get("name")
	price := r.URL.Query().Get("price")
	categoryID := r.URL.Query().Get("category_id")

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if price != "" {
		if val, err := strconv.ParseFloat(price, 64); err == nil {
			query = query.Where("price = ?", val)
		}
	}
	if categoryID != "" {
		if id, err := strconv.Atoi(categoryID); err == nil {
			query = query.Where("category_id = ?", id)
		}
	}

	var products []models.Product
	query.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price <= 0 || p.CategoryID == 0 {
		http.Error(w, "All fields required (name, price, category_id)", http.StatusBadRequest)
		return
	}

	// check category exists
	var cat models.Category
	if err := config.DB.First(&cat, p.CategoryID).Error; err != nil {
		http.Error(w, "Category not found", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&p).Error; err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var updated models.Product
	json.NewDecoder(r.Body).Decode(&updated)

	if updated.Name != "" {
		product.Name = updated.Name
	}
	if updated.Price > 0 {
		product.Price = updated.Price
	}
	if updated.CategoryID > 0 {
		product.CategoryID = updated.CategoryID
	}
	
	config.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Deleted")
}
