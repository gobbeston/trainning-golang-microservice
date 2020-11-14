package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/gobbeston/trainning-golang-microservice/app/utils"
)

func (handler *handler) Health(c *gin.Context) {
	utils.JSONSuccessResponse(c, nil)
}
