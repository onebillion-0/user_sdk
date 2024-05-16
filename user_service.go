package user_sdk

import (
	"github.com/oneliuliu/user_sdk/biz/domain/model"
	"github.com/oneliuliu/user_sdk/biz/infrastrature/db"
)

func CreateUser(req *model.UserInfo) error {
	return db.NewOneBillionDB(db.DB).CreateUser(req)
}
