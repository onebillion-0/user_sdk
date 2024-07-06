package sdk

import (
	"context"

	"github.com/oneliuliu/user_sdk/biz/application/services/school_service"
)

type SchoolLoginController struct {
	service *school_service.LoginService
}

func NewSchoolLoginController(sev *school_service.LoginService) *SchoolLoginController {
	return &SchoolLoginController{service: sev}
}

func (c *SchoolLoginController) Login(ctx context.Context, uid int64, password string) (string, error) {
	token, err := c.service.Login(ctx, uid, password)
	return token, err
}
