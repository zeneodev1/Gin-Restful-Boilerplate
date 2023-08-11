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
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/router"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers/fixtures"
	"gorm.io/gorm"
)

type userControllerTestSuite struct {
	suite.Suite
	router *gin.Engine
	user   *models.User
	tx     *gorm.DB
}

func (s *userControllerTestSuite) SetupSuite() {
	helpers.SetupEnv()
	tx, err := helpers.SetupTx()
	if err != nil {
		s.T().FailNow()
	}
	s.tx = tx
	s.router = router.SetupRouter(tx)
}

func (s *userControllerTestSuite) SetupTest() {
	s.user = fixtures.UserFixture(s.tx)
}

func (s *userControllerTestSuite) TearDownTest() {
	s.tx = helpers.StartOverTx(s.tx)
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
