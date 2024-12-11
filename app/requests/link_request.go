package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LinkRequest struct {
	Name string `valid:"name" json:"name"`
	URL  string `valid:"url" json:"url"`
}

func LinkSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:links,name"},
		"url":  []string{"required", "url"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
			"not_exists:名称已存在",
		},
		"url": []string{
			"required:链接地址为必填项",
			"url:链接地址格式不正确",
		},
	}
	return validate(data, rules, messages)
}
