package contracts

import (
	"project-structure/app/models"
	userRequest "project-structure/app/request/user"
)

type UserServiceInterface interface {
	Create(req userRequest.CreateUserRequest) (models.User, error)
	GetAll() ([]models.User, error)
}
