package user_sdk

import (
	"github.com/oneliuliu/user_sdk/domain/model"
	db2 "github.com/oneliuliu/user_sdk/infrastrature/db"
)

func CreateUser(req *model.UserInfo) error {
	return db2.NewOneBillionDB(db2.DB).CreateUser(req)
}
