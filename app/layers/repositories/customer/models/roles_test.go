package models_test

import (
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CreateRoles(t *testing.T) {
	t.Run("Happy Boom", func(t *testing.T) {
		model := models.Roles{
			Code:        "TEST",
			DisplayName: "Test",
		}

		assert.NotNil(t, model)
	})

	t.Run("Happy - ParseRoleToDB", func(t *testing.T) {
		input := models.Roles{
			Code:        "TEST",
			DisplayName: "Test",
		}

		model, _ := new(models.Roles).ParseRoleToDB(input)

		assert.NotNil(t, model)
	})

	t.Run("Happy - RoleEntity", func(t *testing.T) {
		input := models.Roles{
			Code:        "TEST",
			DisplayName: "Test",
		}

		model, _ := new(models.Roles).ParseRoleToDB(input)
		roleEntity := model.RoleEntity()

		assert.NotNil(t, roleEntity)
		assert.Equal(t, input.Code, roleEntity.Code)
		assert.Equal(t, input.DisplayName, roleEntity.DisplayName)
	})
}