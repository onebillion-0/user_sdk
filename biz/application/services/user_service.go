package services

import (
	"fmt"
	"github.com/oneliuliu/user_sdk/biz/domain/entity"
	"github.com/oneliuliu/user_sdk/biz/domain/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (self *UserService) CreateNewUser(uid int64) error {
	info, err := self.userRepository.FindByID(uid)
	if err != nil {
		return err
	}
	if info != nil {
		return fmt.Errorf("user already exists, uid: %d", uid)
	}
	//todo: add create func
	_, err = self.userRepository.Create(entity.NewUserInfo())
	return err
}

func (self *UserService) GetUserInfo(userId int64) (*entity.UserInfo, error) {
	return self.userRepository.FindByID(userId)
}
