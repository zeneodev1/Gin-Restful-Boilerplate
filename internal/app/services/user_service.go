package services

import (
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/repositories"
)

type UserService interface {
	ListUsers() ([]models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService() UserService {
	return &userService{
		repo: repositories.NewUserRepo(),
	}
}

func (s *userService) ListUsers() ([]models.User, error) {
	return s.repo.ListUsers()
}

func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(id int, user *models.User) error {
	if _, err := s.GetUser(id); err != nil {
		return err
	}

	user.ID = id
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	user, err := s.GetUser(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(user)
}
