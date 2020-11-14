package customer_test

import (
	"github.com/gin-gonic/gin"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	customerHandler "github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/customer"
	customerMock "github.com/gobbeston/trainning-golang-microservice/app/mocks/customer"
	testhelper "github.com/gobbeston/trainning-golang-microservice/app/test_helpers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerMockUseCase := customerMock.NewMockUseCase(ctrl)

	executeWithRequest := func(mockUseCase *customerMock.MockUseCase, jsonRequestBody []byte) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)
		requestURL := "/v1/roles"
		httpRequest, _ := http.NewRequest("POST", requestURL, strings.NewReader(string(jsonRequestBody)))

		customerHandler.NewEndpointHttpHandler(ginEngine, mockUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Success", func(t *testing.T) {
		jsonRequestBody := []byte(`{
									"code": "Boom",
									"displayName": "BoomDisplay"
								}`)

		roleMock := testhelper.GetMockRoleEntity()

		customerMockUseCase.EXPECT().
			CreateRole(gomock.Any()).
			Return(&roleMock, nil)

		res := executeWithRequest(customerMockUseCase, jsonRequestBody)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Failure: CreateRole Throw Error", func(t *testing.T) {
		jsonRequestBody := []byte(`{
									"code": "Boom",
									"displayName": "BoomDisplay"
								}`)

		customerMockUseCase.EXPECT().
			CreateRole(gomock.Any()).
			Return(nil, errors.UnprocessableEntity{Message: "error here"})

		res := executeWithRequest(customerMockUseCase, jsonRequestBody)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Failure : Parse Error", func(t *testing.T) {
		jsonRequestBody := []byte(`{
									"code": "fake-code"
								}`)

		res := executeWithRequest(customerMockUseCase, jsonRequestBody)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
