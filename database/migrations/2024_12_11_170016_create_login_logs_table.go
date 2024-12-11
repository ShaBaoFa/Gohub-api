package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/logger"
	"gohub/pkg/migrate"
	"time"

	"gorm.io/gorm"
)

func init() {

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

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&LoginLog{})
		if err != nil {
			logger.ErrorString("migration", "迁移失败", err.Error())
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&LoginLog{})
		if err != nil {
			logger.ErrorString("migration", "迁移失败", err.Error())
			return
		}
	}

	migrate.Add("2024_12_11_170016_create_login_logs_table", up, down)
}
