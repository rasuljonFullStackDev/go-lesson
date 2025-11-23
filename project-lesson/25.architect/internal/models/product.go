package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Photo string  `json:"photo"`
	Manual string `json:"manual"`
}

// FileFields - modelda qaysi ustunlar faylga tegishli
func (Product) FileFields() map[string][]string {
	return map[string][]string{
		"Photo":  {"jpg", "jpeg", "png", "webp"},
		"Manual": {"pdf"},
	}
}

// MaxFileSize - maksimal ruxsat etilgan hajm (baytda)
func (Product) MaxFileSize() int64 {
	return 5 * 1024 * 1024 // 5 MB
}
