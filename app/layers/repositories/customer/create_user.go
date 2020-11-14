package customer

import (
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer/models"
)

func (r *repo) CreateUser(input *entities.Users) (*entities.Users, error) {
	model, _ := new(models.Users).ParseUserToDB(input)

	if err := r.Conn.Save(&model).Error; err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	return model.UserEntity(), nil
}
