// package models 模型通用属性和方法
package models

import (
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// BaseModel: model base class
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampField : timestamp
type CommonTimestampField struct {
	CreateAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// 在模型中定义 BeforeCreate 方法，确保在创建记录前更新 CreatedAt 字段
func (ctf *CommonTimestampField) BeforeCreate(tx *gorm.DB) (err error) {
    ctf.CreateAt = time.Now()
    return nil
}

// 在模型中定义 BeforeSave 方法，确保在保存记录前更新 UpdatedAt 字段
func (ctf *CommonTimestampField) BeforeSave(tx *gorm.DB) (err error) {
    ctf.UpdatedAt = time.Now()
    return nil
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
    return cast.ToString(a.ID)
}
