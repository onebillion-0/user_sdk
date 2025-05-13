package sdk

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/application/command"

	"github.com/onebillion-0/user_sdk/biz/application/services/school_service"
)

type SchoolLoginController struct {
	service *school_service.LoginService
}

func NewSchoolLoginController(sev *school_service.LoginService) *SchoolLoginController {
	return &SchoolLoginController{service: sev}
}

func (c *SchoolLoginController) Login(ctx context.Context, uid, appid int64, password string) (string, *command.SchoolMemberCommand, error) {
	token, cmd, err := c.service.Login(ctx, uid, appid, password)
	return token, cmd, err
}

type SchoolRegisterController struct {
}
