package user

import (
	"cn.sockstack/gin_demo/dto"
	"cn.sockstack/gin_demo/models"
	"cn.sockstack/gin_demo/pkg/config"
	"fmt"
	"testing"
)

const (
	username = "sockstack"
	password = "123456"
	update_username = "sockstack_update"
)

func init() {
	config.Mysql.Load("../../conf/database.ini").Init()
	models.Init()
}

func TestCreateUser(t *testing.T) {
	userDto := dto.UserDto{
		Username: username,
		Password: password,
	}
	err := CreateUser(userDto)
	if err != nil {
		t.Fail()
	}
}

func TestGetUserByUsername(t *testing.T) {
	user := GetUserByUsername(username)
	if user.Username == "" {
		t.Fail()
	}
	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	user := GetUserByUsername(username)
	userDto := dto.UserDto{
		ID:       user.ID,
		Username: update_username,
		Password: user.Password,
	}
	err := UpdateUser(userDto)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestDeleteUserById(t *testing.T) {
	user := GetUserByUsername(username)
	err := DeleteUserById(user.ID)
	if err != nil {
		t.Fail()
	}
}
