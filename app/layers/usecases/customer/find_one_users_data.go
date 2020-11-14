package customer

import (
	"github.com/gobbeston/trainning-golang-microservice/app/constants"
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
)

func (useCase *useCase) FindOneUserData(input *entities.UsersFilter) (*entities.Users, error) {
	// Find User
	user, err := useCase.CustomerRepo.FindOneUser(input)
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGetDataFromDB}
	}

	// Find Role By RoleID
	filterRole := entities.RolesFilter{
		ID: &user.RoleID,
	}
	role, err := useCase.CustomerRepo.FindOneRole(&filterRole)
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGetDataFromDB}
	}

	user.Roles = role

	return user, nil
}
