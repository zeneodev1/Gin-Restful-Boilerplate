package services

import (
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
)

type UserService struct {
	repo UserRepo
}

type UserRepo interface {
	ListUsers() ([]models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

func NewUserService(repo UserRepo) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) ListUsers() ([]models.User, error) {
	return s.repo.ListUsers()
}

func (s UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s UserService) UpdateUser(id int, user *models.User) error {
	if _, err := s.GetUser(id); err != nil {
		return err
	}

	user.ID = id
	return s.repo.UpdateUser(user)
}

func (s UserService) DeleteUser(id int) error {
	user, err := s.GetUser(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(user)
}
