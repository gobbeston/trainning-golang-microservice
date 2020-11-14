package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/customer/models"
	"github.com/gobbeston/trainning-golang-microservice/app/utils"
)

func (h *handler) CreateRoles(c *gin.Context) {
	//boredom.HandlerInfo(c, nil)

	createRoleRequest, err := new(models.CreateRoleRequestJSON).Parse(c)
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	//entityRole := createRoleRequest.Entity()

	//boredom.FuncDebug(c, h.CustomerUseCase.CreateRole, createRoleRequest)
	roleOutput, err := h.CustomerUseCase.CreateRole(createRoleRequest.Entity())
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	createRole, _ := new(models.CreateRoleResponseJSON).Parse(roleOutput)

	//ginney.JSONSuccessResponse(c, roleOutput)
	utils.JSONSuccessResponse(c, createRole)
}
