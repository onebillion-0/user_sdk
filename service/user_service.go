package service

import (
	model2 "user_sdk/domain/model"
	"user_sdk/infrastrature/db"
)

func CreateUser(req *model2.UserInfo) error {
	return db.NewOneBillionDB(db.DB).CreateUser(req)
}
