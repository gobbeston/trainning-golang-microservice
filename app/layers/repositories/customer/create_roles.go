package customer

import (
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer/models"
)

func (r *repo) CreateRoles(input *entities.Roles) (*entities.Roles, error) {
	model, _ := new(models.Roles).ParseRoleToDB(input)

	if err := r.Conn.Save(&model).Error; err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	return model.RoleEntity(), nil
}
