package repositories

import (
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db}
}

func (r UserRepo) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepo) GetUser(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r UserRepo) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r UserRepo) DeleteUser(user *models.User) error {
	return r.db.Delete(user).Error
}
