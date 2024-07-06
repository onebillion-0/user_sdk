package sdk

import (
	"github.com/onebillion-0/user_sdk/biz/application/services"
	"github.com/onebillion-0/user_sdk/biz/interface/sdk/dto/request"
)

type CreateUserController struct {
	sev *services.UserService
}

func NewCreateUserController(sev *services.UserService) *CreateUserController {
	return &CreateUserController{sev: sev}
}

func (c *CreateUserController) CreateUser(req *request.CreateUserRequest) error {
	cmd, err := req.ToCreateUserCommand()
	if err != nil {
		return err
	}
	return c.sev.CreateNewUser(cmd)
}

func (c *CreateUserController) FindUserByPhoneNumber(phoneNumber string) (bool, error) {
	info, err := c.sev.GetUserInfoByPhoneNumber(phoneNumber)
	if err != nil {
		return false, err
	}
	return info != nil, nil
}
