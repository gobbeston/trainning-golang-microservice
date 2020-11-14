package models_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer/models"
	"testing"
	"time"
)

func TestRepository_CreateUsers(t *testing.T) {

	t.Run("Happy", func(t *testing.T) {
		age := "20"
		birthDate := "1999/09/19"
		address := "92/7"
		phoneNumber := "+66899999999"
		model := models.Users{
			Email:       "tech@mail.co",
			Password:    "P@ssw0rd",
			FirstName:   "dev",
			LastName:    "tech",
			Age:         &age,
			BirthDate:   &birthDate,
			Address:     &address,
			PhoneNumber: &phoneNumber,
			AccessToken: nil,
			Provider:    "OWN",
			LasLogin:    nil,
			StatusID:    1,
			RoleID:      1,
			RoleTypeID:  1,
			CreatedAt:   time.Time{},
			CreatedBy:   "SYSTEM",
			UpdatedAt:   time.Time{},
			UpdatedBy:   "SYSTEM",
			DeletedAt:   nil,
			DeletedBy:   nil,
		}

		assert.NotNil(t, model)
	})

	t.Run("Happy - ParseUserToDB", func(t *testing.T) {
		age := "20"
		birthDate := "1999/09/19"
		address := "92/7"
		phoneNumber := "+66899999999"
		input := entities.Users{
			Email:       "tech@mail.co",
			Password:    "P@ssw0rd",
			FirstName:   "dev",
			LastName:    "tech",
			Age:         &age,
			BirthDate:   &birthDate,
			Address:     &address,
			PhoneNumber: &phoneNumber,
			AccessToken: nil,
			Provider:    "OWN",
			LasLogin:    nil,
			StatusID:    1,
			RoleID:      1,
			RoleTypeID:  1,
			CreatedAt:   time.Time{},
			CreatedBy:   "SYSTEM",
			UpdatedAt:   time.Time{},
			UpdatedBy:   "SYSTEM",
			DeletedAt:   nil,
			DeletedBy:   nil,
		}

		model, _ := new(models.Users).ParseUserToDB(input)

		assert.NotNil(t, model)
	})

	t.Run("Happy - UserEntity", func(t *testing.T) {
		age := "20"
		birthDate := "1999/09/19"
		address := "92/7"
		phoneNumber := "+66899999999"
		input := entities.Users{
			Email:       "tech@mail.co",
			Password:    "P@ssw0rd",
			FirstName:   "dev",
			LastName:    "tech",
			Age:         &age,
			BirthDate:   &birthDate,
			Address:     &address,
			PhoneNumber: &phoneNumber,
			AccessToken: nil,
			Provider:    "OWN",
			LasLogin:    nil,
			StatusID:    1,
			RoleID:      1,
			RoleTypeID:  1,
			CreatedAt:   time.Time{},
			CreatedBy:   "SYSTEM",
			UpdatedAt:   time.Time{},
			UpdatedBy:   "SYSTEM",
			DeletedAt:   nil,
			DeletedBy:   nil,
		}

		model, _ := new(models.Users).ParseUserToDB(input)
		userEntity := model.UserEntity()

		assert.NotNil(t, userEntity)
		assert.Equal(t, input.Email, userEntity.Email)
	})
}
