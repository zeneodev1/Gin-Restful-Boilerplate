package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
	test_setup "github.com/zeneodev1/gin-restful-boilerplate/test"
	"github.com/zeneodev1/gin-restful-boilerplate/test/fixtures"
)

type userControllerTestSuite struct {
	suite.Suite
	router *gin.Engine
	user   *models.User
}

func (s *userControllerTestSuite) SetupSuite() {
	test_setup.SetupRepo()
	s.router = app.SetupRouter()
}

func (s *userControllerTestSuite) SetupTest() {
	s.user = fixtures.UserFixture()
}

func (s *userControllerTestSuite) TearDownTest() {
	fixtures.ClearUsers()
}

func (s *userControllerTestSuite) TestIndex() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Result().StatusCode)

	var result struct {
		Users []models.User `mapped:"users"`
	}
	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Len(s.T(), result.Users, 1)
	user := result.Users[0]
	assert.Equal(s.T(), user.ID, s.user.ID)
	assert.Equal(s.T(), user.FullName, s.user.FullName)
	assert.Equal(s.T(), user.Email, s.user.Email)

}

func (s *userControllerTestSuite) TestShow() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", s.user.ID), nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Result().StatusCode)

	var result struct {
		User *models.User
	}
	json.Unmarshal(w.Body.Bytes(), &result)
	user := result.User

	assert.Equal(s.T(), user.ID, s.user.ID)
	assert.Equal(s.T(), user.FullName, s.user.FullName)
	assert.Equal(s.T(), user.Email, s.user.Email)
}

func TestUserControllerTestTestSuite(t *testing.T) {
	suite.Run(t, new(userControllerTestSuite))
}