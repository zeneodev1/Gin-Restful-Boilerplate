package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/repositories"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers/fixtures"
)

type userRepositoryTestSuite struct {
	suite.Suite
	repo repositories.UserRepo
	user *models.User
}

func (s *userRepositoryTestSuite) SetupSuite() {
	helpers.SetupDB()
	s.repo = repositories.NewUserRepo()
}

func (s *userRepositoryTestSuite) SetupTest() {
	s.user = fixtures.UserFixture()
}

func (s *userRepositoryTestSuite) TearDownTest() {
	fixtures.ClearUsers()
}

func (s *userRepositoryTestSuite) TestListUsers() {
	users, err := s.repo.ListUsers()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
}

func (s *userRepositoryTestSuite) TestGetUser() {
	result, err := s.repo.GetUser(s.user.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.user.ID, result.ID)
	assert.Equal(s.T(), s.user.FullName, result.FullName)
	assert.Equal(s.T(), s.user.Email, result.Email)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}
