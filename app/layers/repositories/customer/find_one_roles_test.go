package customer_test

import (
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer"
	testhelper "github.com/gobbeston/trainning-golang-microservice/app/test_helpers"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepo_FindOneRole(t *testing.T) {
	testhelper.InitEnv()

	var (
		id   = uint(1)
		code = "ADMIN"

		queryPattern = `SELECT * FROM "roles"`
		queryStub    = []map[string]interface{}{{
			"id":         id,
			"code":       code,
			"created_at": time.Time{},
		}}

		filters = entities.RolesFilter{
			ID: &id,
		}
	)

	t.Run("Success", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub)

		CustomerRepo := customer.InitRepo(DB)

		results, err := CustomerRepo.FindOneRole(&filters)

		assert.NoError(t, err)
		assert.Equal(t, id, results.ID)
		assert.Equal(t, code, results.Code)
	})

	t.Run("Failure : got an internal error", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub).WithError(errors.New("error"))
		CustomerRepo := customer.InitRepo(DB)

		result, err := CustomerRepo.FindOneRole(&filters)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
