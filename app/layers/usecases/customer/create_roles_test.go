package customer_test

import (
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	customerUseCase "github.com/gobbeston/trainning-golang-microservice/app/layers/usecases/customer"
	customerMock "github.com/gobbeston/trainning-golang-microservice/app/mocks/customer"
	testhelper "github.com/gobbeston/trainning-golang-microservice/app/test_helpers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCase_CreateRole(t *testing.T) {
	testhelper.InitEnv()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerMockRepo := customerMock.NewMockRepo(ctrl)
	expectedCode := "CodeTest"
	expectedDisplayName := "DisplayNameTest"

	t.Run("Happy", func(t *testing.T) {
		roleEntity := testhelper.GetMockRoleEntity()

		customerMockRepo.EXPECT().
			CreateRoles(&roleEntity).
			Return(&roleEntity, nil)

		useCase := customerUseCase.InitUseCase(customerMockRepo)

		role, err := useCase.CreateRole(&roleEntity)

		assert.Nil(t, err)
		assert.Equal(t, expectedCode, role.Code)
		assert.Equal(t, expectedDisplayName, role.DisplayName)
	})

	t.Run("Fail: repo return error", func(t *testing.T) {
		roleEntity := testhelper.GetMockRoleEntity()

		customerMockRepo.EXPECT().
			CreateRoles(&roleEntity).
			Return(nil, errors.InternalError{Message: "error here"})

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		_, err := useCase.CreateRole(&roleEntity)
		assert.Error(t, err)
	})
}
