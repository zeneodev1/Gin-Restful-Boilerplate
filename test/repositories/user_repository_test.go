package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/repositories"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers"
	"github.com/zeneodev1/gin-restful-boilerplate/test/helpers/fixtures"
	"gorm.io/gorm"
)

type userRepositoryTestSuite struct {
	suite.Suite
	repo repositories.UserRepo
	user *models.User
	tx   *gorm.DB
}

func (s *userRepositoryTestSuite) SetupSuite() {
	helpers.SetupEnv()
	tx, err := helpers.SetupTx()
	if err != nil {
		s.T().FailNow()
	}
	s.tx = tx
	s.repo = repositories.NewUserRepo(tx)
}

func (s *userRepositoryTestSuite) SetupTest() {
	s.user = fixtures.UserFixture(s.tx)
}

func (s *userRepositoryTestSuite) TearDownTest() {
	s.tx = helpers.StartOverTx(s.tx)
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
