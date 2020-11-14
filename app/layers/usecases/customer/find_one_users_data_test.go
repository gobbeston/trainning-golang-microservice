package customer_test

import (
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	customerUseCase "github.com/gobbeston/trainning-golang-microservice/app/layers/usecases/customer"
	customerMock "github.com/gobbeston/trainning-golang-microservice/app/mocks/customer"
	testhelper "github.com/gobbeston/trainning-golang-microservice/app/test_helpers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCase_FindOneUserData(t *testing.T) {
	testhelper.InitEnv()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerMockRepo := customerMock.NewMockRepo(ctrl)

	var (
		expectedUserId = uint(1)
		expectedEmail  = "tech@mail.co"
		expectedRoleId = uint(1)
	)

	t.Run("Happy", func(t *testing.T) {
		userEntity := testhelper.GetMockUserEntity()
		roleEntity := testhelper.GetMockRoleEntity()

		userFilter := entities.UsersFilter{
			ID:    &expectedUserId,
			Email: &expectedEmail,
		}
		customerMockRepo.EXPECT().
			FindOneUser(&userFilter).
			Return(&userEntity, nil)

		roleFilter := entities.RolesFilter{
			ID: &expectedRoleId,
		}
		customerMockRepo.EXPECT().
			FindOneRole(&roleFilter).
			Return(&roleEntity, nil)

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		user, err := useCase.FindOneUserData(&userFilter)
		assert.Nil(t, err)
		assert.Equal(t, expectedUserId, user.ID)
		assert.Equal(t, expectedEmail, user.Email)
		assert.Equal(t, expectedRoleId, user.Roles.ID)
	})

	t.Run("Fail: repo FindOneUser return error", func(t *testing.T) {
		userFilter := entities.UsersFilter{
			ID:    &expectedUserId,
			Email: &expectedEmail,
		}
		customerMockRepo.EXPECT().
			FindOneUser(&userFilter).
			Return(nil, errors.InternalError{Message: "error here"})

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		_, err := useCase.FindOneUserData(&userFilter)
		assert.Error(t, err)
	})

	t.Run("Fail: repo FindOneRole return error", func(t *testing.T) {
		userEntity := testhelper.GetMockUserEntity()

		userFilter := entities.UsersFilter{
			ID:    &expectedUserId,
			Email: &expectedEmail,
		}
		customerMockRepo.EXPECT().
			FindOneUser(&userFilter).
			Return(&userEntity, nil)

		roleFilter := entities.RolesFilter{
			ID: &expectedRoleId,
		}
		customerMockRepo.EXPECT().
			FindOneRole(&roleFilter).
			Return(nil, errors.InternalError{Message: "error here"})

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		_, err := useCase.FindOneUserData(&userFilter)
		assert.Error(t, err)
	})
}
