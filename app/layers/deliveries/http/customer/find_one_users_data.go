package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/customer/models"
	"github.com/gobbeston/trainning-golang-microservice/app/utils"
)

func (h *handler) FindOneUserData(c *gin.Context) {
	//boredom.HandlerInfo(c, models.FindOneUserDataRequestJSON{})

	findOneUserDataRequest, err := new(models.FindOneUserDataRequestJSON).Parse(c)
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	//boredom.FuncDebug(c, h.UsersUseCase.FindOneUser, findOneUserRequest)
	UserOutput, err := h.CustomerUseCase.FindOneUserData(findOneUserDataRequest.Entity())

	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	user, _ := new(models.FindOneUserDataResponseJSON).Parse(*UserOutput)

	//ginney.JSONSuccessResponse(c, user)
	utils.JSONSuccessResponse(c, user)
}
