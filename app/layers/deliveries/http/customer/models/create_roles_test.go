package models_test

import (
	"fmt"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/customer/models"
	testhelper "github.com/gobbeston/trainning-golang-microservice/app/test_helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRoleRequestJSON_Parse(t *testing.T) {
	var (
		code        = "BOOM"
		displayName = "BoomDisplay"
	)

	t.Run("Success", func(t *testing.T) {
		params := fmt.Sprintf(`{
			"code": "%s",
			"displayName": "%s"
		}`,
			code,
			displayName,
		)

		ctx := testhelper.MakeStubContext("POST", "/", params)

		createRoleRequestJSON, err := new(models.CreateRoleRequestJSON).Parse(ctx)

		assert.Nil(t, err)
		assert.NotEmpty(t, createRoleRequestJSON)
		assert.Equal(t, code, createRoleRequestJSON.Code)
		assert.Equal(t, displayName, createRoleRequestJSON.DisplayName)
	})

	t.Run("Failure", func(t *testing.T) {
		emptyParams := ""

		ctx := testhelper.MakeStubContext("POST", "/", emptyParams)

		createRequestJSON, err := new(models.CreateRoleRequestJSON).Parse(ctx)

		assert.Error(t, err)
		assert.Nil(t, createRequestJSON)
	})
}

func TestCreateRoleRequestJSON_Entity(t *testing.T) {
	var (
		code        = "BOOM"
		displayName = "BoomDisplay"
	)

	t.Run("Success", func(t *testing.T) {
		createRoleRequestJSON := models.CreateRoleRequestJSON{
			Code:        "BOOM",
			DisplayName: "BoomDisplay",
		}

		role := createRoleRequestJSON.Entity()

		assert.NotEmpty(t, role)
		assert.Equal(t, code, role.Code)
		assert.Equal(t, displayName, role.DisplayName)
	})
}

func TestCreateRoleResponseJSON_Parse(t *testing.T) {
	var (
		code        = "BOOM"
		displayName = "BoomDisplay"
	)

	t.Run("Success", func(t *testing.T) {
		params := fmt.Sprintf(`{
			"code": "%s",
			"displayName": "%s"
		}`,
			code,
			displayName,
		)

		ctx := testhelper.MakeStubContext("POST", "/", params)

		createRoleRequestJSON, err := new(models.CreateRoleRequestJSON).Parse(ctx)
		createRoleResponseJSON, err := new(models.CreateRoleResponseJSON).Parse(createRoleRequestJSON)

		assert.Nil(t, err)
		assert.NotEmpty(t, createRoleResponseJSON)
		assert.Equal(t, code, createRoleResponseJSON.Code)
		assert.Equal(t, displayName, createRoleResponseJSON.DisplayName)
	})
}