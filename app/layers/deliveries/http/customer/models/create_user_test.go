package models_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/customer/models"
	testhelper "github.com/gobbeston/trainning-golang-microservice/app/test_helpers"
	"testing"
)

func TestCreateUserRequestJSON_Entity(t *testing.T) {
	var (
		email     = "tech@mail.co"
		password  = "P@ssw0rd"
		firstName = "n"
		lastName  = "digital"
	)

	t.Run("Success", func(t *testing.T) {
		age := "20"
		birthDate := "1999/09/09"
		address := "92/7"
		phoneNumber := "+66899999999"
		createUserRequestJSON := models.CreateUserRequestJSON{
			Email:       "tech@mail.co",
			Password:    "P@ssw0rd",
			FirstName:   "n",
			LastName:    "digital",
			Age:         &age,
			BirthDate:   &birthDate,
			Address:     &address,
			PhoneNumber: &phoneNumber,
			Provider:    "OWN",
			StatusID:    1,
			RoleID:      1,
			RoleTypeID:  1,
			CreatedBy:   "N System",
		}

		user := createUserRequestJSON.Entity()

		assert.NotEmpty(t, user)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, firstName, user.FirstName)
		assert.Equal(t, lastName, user.LastName)
	})
}

func TestCreateUserRequestJSON_Parse(t *testing.T) {
	var (
		email      = "tech@mail.co"
		password   = "P@ssw0rd"
		firstName  = "n"
		lastName   = "digital"
		provider   = "OWN"
		statusID   = uint(1)
		roleID     = uint(1)
		roleTypeID = uint(1)
		createdBy  = "N System"
	)

	t.Run("Success", func(t *testing.T) {
		params := fmt.Sprintf(`{
			"email": "%s",
			"password": "%s",
			"firstName": "%s",
			"lastName": "%s",
			"provider": "%s",
			"statusId": %d,
			"roleId": %d,
			"roleTypeId": %d,
			"createdBy": "%s"
		}`,
			email,
			password,
			firstName,
			lastName,
			provider,
			statusID,
			roleID,
			roleTypeID,
			createdBy,
		)

		ctx := testhelper.MakeStubContext("POST", "/", params)

		createUserRequestJSON, err := new(models.CreateUserRequestJSON).Parse(ctx)

		assert.Nil(t, err)
		assert.NotEmpty(t, createUserRequestJSON)
		assert.Equal(t, email, createUserRequestJSON.Email)
		assert.Equal(t, password, createUserRequestJSON.Password)
		assert.Equal(t, firstName, createUserRequestJSON.FirstName)
		assert.Equal(t, lastName, createUserRequestJSON.LastName)
	})

	t.Run("Failure", func(t *testing.T) {
		emptyParams := ""

		ctx := testhelper.MakeStubContext("POST", "/", emptyParams)

		createRequestJSON, err := new(models.CreateUserRequestJSON).Parse(ctx)

		assert.Error(t, err)
		assert.Nil(t, createRequestJSON)
	})
}
