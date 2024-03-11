// Package user --> user model-related logic
package user

import (
	"goapihub/app/models"
	"goapihub/pkg/database"
)

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Email string `json:"-"`
	Phone string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() error {
    result :=  database.DB.Create(&userModel)
	if result.Error != nil {
		return result.Error
	}
	// success
	return nil
}