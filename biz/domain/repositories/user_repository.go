package repositories

import (
	"github.com/oneliuliu/user_sdk/biz/domain/entity"
)

type UserRepository interface {
	FindByID(uid int64) (*entity.UserInfo, error)
	Create(user *entity.UserInfo) (*entity.UserInfo, error)
	Update(user *entity.UserInfo) (*entity.UserInfo, error)
	Delete(uid int64) error
	FindByPhoneNumber(number string) (*entity.UserInfo, error)
}
