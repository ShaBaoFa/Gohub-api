package login_log

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idstr string) (loginLog LoginLog) {
	database.DB.Where("id", idstr).First(&loginLog)
	return
}

func GetBy(field, value string) (loginLog LoginLog) {
	database.DB.Where("? = ?", field, value).First(&loginLog)
	return
}

func All() (loginLogs []LoginLog) {
	database.DB.Find(&loginLogs)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(LoginLog{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (loginLogs []LoginLog, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(LoginLog{}),
		&loginLogs,
		app.V1URL(database.TableName(&LoginLog{})),
		perPage,
	)
	return
}
