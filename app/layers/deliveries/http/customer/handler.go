package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/usecases/customer"
)

type handler struct {
	CustomerUseCase customer.UseCase
}

// NewEndpointHandler routing
func NewEndpointHttpHandler(ginEngine *gin.Engine, customerUseCase customer.UseCase) {
	handler := &handler{
		CustomerUseCase: customerUseCase,
	}

	v1 := ginEngine.Group("v1")
	{
		v1.POST("/users", handler.CreateUser)
		v1.POST("/roles", handler.CreateRoles)
	}
}
