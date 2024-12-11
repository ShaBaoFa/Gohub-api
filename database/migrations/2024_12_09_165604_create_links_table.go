package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/logger"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Link{})
		if err != nil {
			logger.ErrorString("migration", "迁移失败", err.Error())
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Link{})
		if err != nil {
			logger.ErrorString("migration", "迁移失败", err.Error())
			return
		}
	}

	migrate.Add("2024_12_09_165604_create_links_table", up, down)
}
