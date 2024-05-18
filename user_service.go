package user_sdk

import (
	"errors"
	"github.com/oneliuliu/user_sdk/biz/application/services"
	"github.com/oneliuliu/user_sdk/biz/infrastrature/db/mysql"
	"github.com/oneliuliu/user_sdk/biz/interface/sdk"
	"github.com/oneliuliu/user_sdk/biz/interface/sdk/dto/request"
)

func CreateUser(req *request.CreateUserRequest) error {
	if req == nil {
		return errors.New("req is nil")
	}
	gormRepo := mysql.NewGormUserRepository(mysql.Connection)
	userService := services.NewUserService(gormRepo)
	return sdk.NewCreateUserController(userService).CreateUser(req)
}
