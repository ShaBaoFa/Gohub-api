package models

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time      `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at,omitempty"` // 删除时间
	Remark    string         `gorm:"column:remark;comment:备注" json:"remark"`                     // 备注
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
