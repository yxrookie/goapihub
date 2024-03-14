// Package user --> user model-related logic
package user

import (
	"goapihub/app/models"
	"goapihub/pkg/database"
	"goapihub/pkg/hash"
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

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
    return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
    result := database.DB.Save(&userModel)
    return result.RowsAffected
}