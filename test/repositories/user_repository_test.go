package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/repositories"
	test_setup "github.com/zeneodev1/gin-restful-boilerplate/test"
	"github.com/zeneodev1/gin-restful-boilerplate/test/fixtures"
)

type userRepositoryTestSuite struct {
	suite.Suite
	repo repositories.UserRepo
	user *models.User
}

func (s *userRepositoryTestSuite) SetupSuite() {
	test_setup.SetupRepo()
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
