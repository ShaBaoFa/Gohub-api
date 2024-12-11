package v1

import (
	"gohub/app/models/login_log"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginLogsController struct {
	BaseAPIController
}

func (ctrl *LoginLogsController) Index(c *gin.Context) {
	loginLogs := login_log.All()
	response.Data(c, loginLogs)
}
