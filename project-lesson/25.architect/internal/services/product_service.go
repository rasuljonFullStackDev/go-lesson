package services

import "25.architect/internal/models"

type ProductService struct {
	CrudService[models.Product]
}
