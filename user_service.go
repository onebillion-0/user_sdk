package user_sdk

import (
	"errors"
	"github.com/onebillion-0/user_sdk/biz/application/services"
	"github.com/onebillion-0/user_sdk/biz/infrastrature/db/mysql"
	"github.com/onebillion-0/user_sdk/biz/interface/sdk"
	"github.com/onebillion-0/user_sdk/biz/interface/sdk/dto/request"
)

func CreateUser(req *request.CreateUserRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}
	gormRepo := mysql.NewGormUserRepository(mysql.Connection)
	userService := services.NewUserService(gormRepo)
	return sdk.NewCreateUserController(userService).CreateUser(req)
}

func CheckUserRegisteredByPhoneNumber(phoneNumber string) (bool, error) {
	gormRepo := mysql.NewGormUserRepository(mysql.Connection)
	userService := services.NewUserService(gormRepo)
	return sdk.NewCreateUserController(userService).FindUserByPhoneNumber(phoneNumber)
}
