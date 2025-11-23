package handlers

import (
	"encoding/json"
	"11.api-crud/config"
	"11.api-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	config.DB.Preload("Products").Find(&categories)
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if c.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}
	if err := config.DB.Create(&c).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var c models.Category
	if err := config.DB.First(&c, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	var updated models.Category
	json.NewDecoder(r.Body).Decode(&updated)
	if updated.Name != "" {
		c.Name = updated.Name
	}
	config.DB.Save(&c)
	json.NewEncoder(w).Encode(c)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	config.DB.Delete(&models.Category{}, id)
	w.WriteHeader(http.StatusNoContent)
}
