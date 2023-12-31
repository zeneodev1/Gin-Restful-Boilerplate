package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
)

type UserCtrl struct {
	userService UserService
}

type UserService interface {
	ListUsers() ([]models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}

func NewUserCtrl(userService UserService) UserCtrl {
	return UserCtrl{
		userService: userService,
	}
}

func (ctrl UserCtrl) Index(c *gin.Context) {
	users, err := ctrl.userService.ListUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (ctrl UserCtrl) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctrl.userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (ctrl UserCtrl) Create(c *gin.Context) {
	var t models.User
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.userService.CreateUser(&t); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl UserCtrl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var t models.User
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.userService.UpdateUser(id, &t); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": t})
}

func (ctrl UserCtrl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
