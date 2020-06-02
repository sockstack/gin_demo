package user

import (
	"cn.sockstack/gin_demo/dto"
	"cn.sockstack/gin_demo/models"
)

//CreateUser 创建用户
func CreateUser(dto dto.UserDto) error {
	user := models.User{}
	user.Username = dto.Username
	user.Password = dto.Password

	err := models.DB.Create(&user).Error

	return err
}

//GetUserByUsername 通过用户名查询用户
func GetUserByUsername(username string) models.User {
	user := models.User{}
	models.DB.Find(&user, models.DB.Where("username = ?",username))

	return user
}

//UpdateUser 更新用户信息
func UpdateUser(userDto dto.UserDto) error {
	user := models.User{}
	if userDto.Username != "" {
		user.Username = userDto.Username
	}

	if userDto.Password != "" {
		user.Password = userDto.Password
	}

	err := models.DB.Model(&user).Update(&user).Error

	return err
}

//DeleteUserById 通过Id删除用户
func DeleteUserById(id uint) error {
	err := models.DB.Delete(&models.User{ID: id}).Error
	return err
}
