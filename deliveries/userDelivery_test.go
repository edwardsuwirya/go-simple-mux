package deliveries

import (
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

type serviceMock struct {
	mock.Mock
}

func (s serviceMock) Register(newUser *models.User) error {
	return nil
}

func (s serviceMock) GetUserInfo(id string) *models.User {
	return &models.User{
		Id:        "123",
		FirstName: "Dummy First Name",
		LastName:  "Dummy Last Name",
		IsActive:  "T",
	}
}

func (s serviceMock) Unregister(id string) error {
	return nil
}

func (s serviceMock) UpdateInfo(id string, newUser *models.User) error {
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
	assert.Equal(suite.T(), "123", resultTest.(*UserDelivery).service.GetUserInfo("").Id)
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
	resultTest := NewUserDelivery(suite.routerTest, suite.parserTest, suite.responderTest, suite.serviceTest)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dummy-user?id=123", nil)

	handler := resultTest.(*UserDelivery).userRoute
	suite.routerTest.HandleFunc("/dummy-user", handler)
	suite.routerTest.ServeHTTP(rr, req)
	respTest := rr.Result()
	respBody := new(appHttpResponse.Response)
	if err := json.NewDecoder(respTest.Body).Decode(respBody); err != nil {
	}
	assert.Equal(suite.T(), rr.Code, 200)
	assert.Equal(suite.T(), respBody.Data.(map[string]interface{})["firstName"], "Dummy First Name")
}

func TestUserDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(UserDeliveryTestSuite))
}
