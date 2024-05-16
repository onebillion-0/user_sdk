package db

import (
	"context"
	"fmt"
	"github.com/oneliuliu/user_sdk/biz/domain/model"
	"gorm.io/gorm"
)

type OneBillionDB struct {
	db *gorm.DB
}

func NewOneBillionDB(db *gorm.DB) *OneBillionDB {

	return &OneBillionDB{db: db}
}

func (o *OneBillionDB) CreateUser(info *model.UserInfo) error {
	err := o.db.AutoMigrate(&model.UserInfo{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	result := o.db.Create(info)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return result.Error
}

func (o *OneBillionDB) GetUserById(ctx context.Context, userId int64) (*model.UserInfo, error) {
	return nil, nil
}
