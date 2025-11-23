package user

import (
	"project-structure/app/models"
	userRequest "project-structure/app/request/user"
	"project-structure/app/resource/services/api/user/contracts"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) contracts.UserServiceInterface {
	return &UserService{DB: db}
}

func (s *UserService) Create(req userRequest.CreateUserRequest) (models.User, error) {
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // haqiqiy loyihada hashing kerak
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
