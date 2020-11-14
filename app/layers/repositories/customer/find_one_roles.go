package customer

import (
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/errors"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer/models"
)

func (r *repo) FindOneRole(filter *entities.RolesFilter) (*entities.Roles, error) {
	var result models.Roles

	if err := r.Conn.Where(makeRoleFilterAttr(*filter)).Find(&result).Error; err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	output := result.RoleEntity()

	return output, nil
}

func makeRoleFilterAttr(filter entities.RolesFilter) (roleFilterAttr models.Roles) {
	roleFilterAttr = models.Roles{}
	if filter.ID != nil {
		roleFilterAttr.ID = *filter.ID
	}

	return roleFilterAttr
}
