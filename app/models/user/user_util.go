package user

import "goapihub/pkg/database"

// IsEmailExist: judge whether emial is resgitered
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist: judge whether phone number is registered
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone: get user by phone 
func GetByPhone(phone string)(userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti: get user by phone or email or username
func GetByMulti(loginID string)(userModel User) {
	database.DB.Where("phone = ?", loginID).
	Or("email = ?", loginID).
	Or("name = ?", loginID).
	First(&userModel)
	return
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
    database.DB.Where("id", idstr).First(&userModel)
    return
}


// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
    database.DB.Where("email = ?", email).First(&userModel)
    return
}