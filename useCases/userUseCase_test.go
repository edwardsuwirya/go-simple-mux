package useCases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gosimplemux/models"
	"gosimplemux/repositories"
	"testing"
)

var dummyUsers = []*models.User{
	&models.User{
		Id:        "1",
		FirstName: "Dummy First Name 1",
		LastName:  "Dummy Last Name 1",
		IsActive:  "A",
	}, &models.User{
		Id:        "2",
		FirstName: "Dummy First Name 2",
		LastName:  "Dummy Last Name 2",
		IsActive:  "A",
	},
}

type repoMock struct {
	mock.Mock
}

func (r *repoMock) FindOneById(id string) (*models.User, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), nil
	}
	return nil, args.Get(1).(error)
}

func (r *repoMock) Create(newUser *models.User) error {
	args := r.Called(newUser)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMock) Update(id string, newUser *models.User) error {
	panic("implement me")
}

func (r *repoMock) Delete(id string) error {
	panic("implement me")
}

type UserUseCaseTestSuite struct {
	suite.Suite
	repoTest repositories.UserRepository
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.repoTest = new(repoMock)
}

func (suite *UserUseCaseTestSuite) TestBuildNewUserUseCase() {
	useCaseTest := NewUserUseCase(suite.repoTest)
	assert.NotNil(suite.T(), useCaseTest)
	var dummyImpl *IUserUseCase
	assert.Implements(suite.T(), dummyImpl, useCaseTest)
}

func (suite *UserUseCaseTestSuite) TestRegister() {
	suite.repoTest.(*repoMock).On("Create", dummyUsers[0]).Return(nil)
	useCaseTest := NewUserUseCase(suite.repoTest)
	err := useCaseTest.Register(dummyUsers[0])
	assert.Nil(suite.T(), err)
}
func (suite *UserUseCaseTestSuite) TestFailedRegister() {
	suite.repoTest.(*repoMock).On("Create", dummyUsers[0]).Return(errors.New("Failed"))
	useCaseTest := NewUserUseCase(suite.repoTest)
	err := useCaseTest.Register(dummyUsers[0])
	assert.NotNil(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestGetUserInfo() {
	suite.repoTest.(*repoMock).On("FindOneById", dummyUsers[0].Id).Return(dummyUsers[0], nil)
	useCaseTest := NewUserUseCase(suite.repoTest)
	dummyUser := useCaseTest.GetUserInfo(dummyUsers[0].Id)
	assert.Equal(suite.T(), dummyUsers[0].Id, dummyUser.Id)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
