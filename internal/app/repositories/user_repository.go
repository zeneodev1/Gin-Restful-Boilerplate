package repositories

import (
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
)

type UserRepo interface {
	ListUsers() ([]models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	DeleteUser(*models.User) error
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := Repo.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetUser(id int) (*models.User, error) {
	var user models.User
	if err := Repo.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) CreateUser(user *models.User) error {
	return Repo.Create(user).Error
}

func (r *userRepo) UpdateUser(user *models.User) error {
	return Repo.Save(user).Error
}

func (r *userRepo) DeleteUser(user *models.User) error {
	return Repo.Delete(user).Error
}
