package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/repositories"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/services"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers/fixtures"
	"gorm.io/gorm"
)

type userServiceTestSuite struct {
	suite.Suite
	service services.UserService
	user    *models.User
	tx      *gorm.DB
}

func (s *userServiceTestSuite) SetupSuite() {
	helpers.SetupEnv()
	tx, err := helpers.SetupTx()
	if err != nil {
		s.T().FailNow()
	}
	repo := repositories.NewUserRepo(tx)
	s.service = services.NewUserService(repo)
	s.tx = tx
}

func (s *userServiceTestSuite) SetupTest() {
	s.user = fixtures.UserFixture(s.tx)
}

func (s *userServiceTestSuite) TearDownTest() {
	s.tx = helpers.StartOverTx(s.tx)
}

func (s *userServiceTestSuite) TearDownSuite() {
	helpers.Rollback(s.tx)
}

func (s *userServiceTestSuite) TestListUsers() {
	users, err := s.service.ListUsers()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	user := users[0]
	assert.Equal(s.T(), user.ID, s.user.ID)
	assert.Equal(s.T(), user.FullName, s.user.FullName)
	assert.Equal(s.T(), user.Email, s.user.Email)
}

func (s *userServiceTestSuite) TestGetUser() {
	user, err := s.service.GetUser(s.user.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), user.ID, s.user.ID)
	assert.Equal(s.T(), user.FullName, s.user.FullName)
	assert.Equal(s.T(), user.Email, s.user.Email)

	_, err = s.service.GetUser(4012)
	assert.Error(s.T(), err)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(userServiceTestSuite))
}
