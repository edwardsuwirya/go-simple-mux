package repositories

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gosimplemux/models"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.db = DbPrep()
}

func (suite *UserRepositoryTestSuite) TestBuildNewUserRepository() {
	repoTest := NewUserRepository(suite.db)
	var dummyImpl *UserRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *UserRepositoryTestSuite) TestFindOneById() {
	repoTest := NewUserRepository(suite.db)
	userDummy, err := repoTest.FindOneById("1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "1", userDummy.Id)
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	repoTest := NewUserRepository(suite.db)
	dummyUser := &models.User{
		FirstName: "Fn 3",
		LastName:  "Ln 3",
		IsActive:  "A",
	}
	err := repoTest.Create(dummyUser)
	assert.Nil(suite.T(), err)
	userDummy, err := repoTest.FindOneById(dummyUser.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Fn 3", userDummy.FirstName)
}
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
