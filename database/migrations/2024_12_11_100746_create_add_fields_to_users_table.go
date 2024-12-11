package migrations

import (
	"database/sql"
	"gohub/pkg/logger"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(10);"`
		Introduction string `gorm:"type:varchar(255);"`
		Avatar       string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&User{})
		if err != nil {
			logger.ErrorString("migration", "迁移失败", err.Error())
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropColumn(&User{}, "City")
		if err != nil {
			return
		}
		err1 := migrator.DropColumn(&User{}, "Introduction")
		if err1 != nil {
			return
		}
		err2 := migrator.DropColumn(&User{}, "Avatar")
		if err2 != nil {
			return
		}
	}

	migrate.Add("2024_12_11_100746_create_add_fields_to_users_table", up, down)
}
