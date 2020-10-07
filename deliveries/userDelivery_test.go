package deliveries

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gosimplemux/appUtils/appHttpParser"
	"gosimplemux/appUtils/appHttpResponse"
	"gosimplemux/models"
	"gosimplemux/useCases"
	"net/http"
	"net/http/httptest"
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

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Register(newUser *models.User) error {
	return nil
}

func (s *serviceMock) GetUserInfo(id string) *models.User {
	args := s.Called(id)
	if args.Get(0) != nil {
		userRes := args[0].(*models.User)
		return userRes
	}
	return nil
}

func (s *serviceMock) Unregister(id string) error {
	return nil
}

func (s *serviceMock) UpdateInfo(id string, newUser *models.User) error {
	return nil
}

type UserDeliveryTestSuite struct {
	suite.Suite
	routerTest    *mux.Router
	parserTest    *appHttpParser.JsonParser
	responderTest appHttpResponse.IResponder
	serviceTest   useCases.IUserUseCase
}

func (suite *UserDeliveryTestSuite) SetupTest() {
	suite.routerTest = mux.NewRouter()
	suite.parserTest = appHttpParser.NewJsonParser()
	suite.responderTest = appHttpResponse.NewJSONResponder()
	suite.serviceTest = new(serviceMock)
}

func (suite *UserDeliveryTestSuite) TestBuildNewUserDelivery() {
	resultTest := NewUserDelivery(suite.routerTest, suite.parserTest, suite.responderTest, suite.serviceTest)
	var dummyImpl *IDelivery
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	assert.NotNil(suite.T(), resultTest.(*UserDelivery).service)
}

func (suite *UserDeliveryTestSuite) TestInitRoute() {
	resultTest := NewUserDelivery(suite.routerTest, suite.parserTest, suite.responderTest, suite.serviceTest)
	resultTest.InitRoute()

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user", nil)
	suite.routerTest.ServeHTTP(rr, req)
	assert.Equal(suite.T(), rr.Code, 200)

	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user", nil)
	suite.routerTest.ServeHTTP(rr, req)
	assert.Equal(suite.T(), rr.Code, 200)

	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user-wrong-path", nil)
	suite.routerTest.ServeHTTP(rr, req)
	assert.Equal(suite.T(), rr.Code, 404)
}

func (suite *UserDeliveryTestSuite) TestUserRoute() {
	//Setup usecase method yang sudah kita kondisikan
	dummyId := dummyUsers[1].Id
	suite.serviceTest.(*serviceMock).On("GetUserInfo", dummyId).Return(dummyUsers[1])

	//Buat NewUserDelivery function dummy
	resultTest := NewUserDelivery(suite.routerTest, suite.parserTest, suite.responderTest, suite.serviceTest)

	//Buat httptest response dummy
	rr := httptest.NewRecorder()
	//Buat httptest request dummy
	req, _ := http.NewRequest("GET", "/dummy-user?id="+dummyId, nil)

	//Register handler function yang mau kita test
	handler := resultTest.(*UserDelivery).userRoute
	suite.routerTest.HandleFunc("/dummy-user", handler)

	//Naikan serverHTTP nya
	suite.routerTest.ServeHTTP(rr, req)

	//Cek hasil
	respTest := rr.Result()
	respBody := new(appHttpResponse.Response)
	if err := json.NewDecoder(respTest.Body).Decode(respBody); err != nil {
	}

	//Pengujian ekspetasi
	assert.Equal(suite.T(), rr.Code, 200)
	assert.Equal(suite.T(), respBody.Data.(map[string]interface{})["firstName"], dummyUsers[1].FirstName)
}

func (suite *UserDeliveryTestSuite) TestUserPostRoute() {
	resultTest := NewUserDelivery(suite.routerTest, suite.parserTest, suite.responderTest, suite.serviceTest)
	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyUsers[0])
	handler := resultTest.(*UserDelivery).userPostRoute
	suite.routerTest.HandleFunc("/dummy-user", handler)
	req, _ := http.NewRequest("POST", "/dummy-user", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, req)
	respTest := rr.Result()
	respBody := new(appHttpResponse.Response)

	if err := json.NewDecoder(respTest.Body).Decode(respBody); err != nil {
	}
	assert.Equal(suite.T(), rr.Code, 200)
	assert.Equal(suite.T(), respBody.Data.(map[string]interface{})["firstName"], "Dummy First Name 1")
}

func TestUserDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(UserDeliveryTestSuite))
}
