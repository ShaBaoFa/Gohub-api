// Package login_log 模型
package login_log

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"time"
)

type LoginLog struct {
	models.BaseModel
	Name       string    `gorm:"column:name;not null;comment:用户名" json:"name"`                          // 用户名
	IP         string    `gorm:"column:ip;comment:登录IP地址" json:"ip"`                                    // 登录IP地址
	IPLocation string    `gorm:"column:ip_location;comment:IP所属地" json:"ip_location"`                   // IP所属地
	Os         string    `gorm:"column:os;comment:操作系统" json:"os"`                                      // 操作系统
	Browser    string    `gorm:"column:browser;comment:浏览器" json:"browser"`                             // 浏览器
	Status     int32     `gorm:"column:status;not null;default:1;comment:登录状态 (1成功 2失败)" json:"status"` // 登录状态 (1成功 2失败)
	Message    string    `gorm:"column:message;comment:提示消息" json:"message"`                            // 提示消息
	LoginTime  time.Time `gorm:"column:login_time;not null;comment:登录时间" json:"login_time"`             // 登录时间
	// Put fields in here

	models.CommonTimestampsField
}

func (loginLog *LoginLog) Create() {
	database.DB.Create(&loginLog)
}

func (loginLog *LoginLog) Save() (rowsAffected int64) {
	result := database.DB.Save(&loginLog)
	return result.RowsAffected
}

func (loginLog *LoginLog) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&loginLog)
	return result.RowsAffected
}
