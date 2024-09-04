package user

import "github.com/cyansobble/global"

func GetUserByID(id uint) (User, error) {
	var user User
	err := global.DB.First(&user, id).Error
	return user, err
}

func SaveUser(u User) error {
	result := global.DB.Save(&u)
	return result.Error
}
