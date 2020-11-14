package customer

import (
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer/models"
)

func (r *repo) FindOneUser(filter *entities.UsersFilter) (*entities.Users, error) {
	var result models.Users

	if err := r.Conn.Where(makeUserFilterAttr(*filter)).Find(&result).Error; err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	output := result.UserEntity()

	return output, nil
}

func makeUserFilterAttr(filter entities.UsersFilter) (userFilterAttr models.Users) {
	userFilterAttr = models.Users{}
	if filter.ID != nil {
		userFilterAttr.ID = *filter.ID
	}

	if filter.Email != nil {
		userFilterAttr.Email = *filter.Email
	}

	return userFilterAttr
}
