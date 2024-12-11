package link

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/cache"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"
	"time"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {
	// 设置缓存 key
	cacheKey := "links:all"
	// 设置过期时间
	cacheExpiration := 30 * time.Minute
	// 取缓存
	cache.GetObject(cacheKey, &links)
	// 如果数据为空
	if helpers.Empty(links) {
		// 查数据库
		links = All()
		// 如果link依旧为空,直接返回
		if helpers.Empty(links) {
			return
		}
		// 存入缓存
		cache.Set(cacheKey, links, cacheExpiration)
	}
	return
}
