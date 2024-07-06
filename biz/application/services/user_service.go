package services

import (
	"fmt"
	"github.com/onebillion-0/user_sdk/biz/application/command"
	"github.com/onebillion-0/user_sdk/biz/domain/entity"
	"github.com/onebillion-0/user_sdk/biz/domain/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateNewUser(userCommand *command.CreateUserCommand) error {
	info, err := s.userRepository.FindByID(userCommand.Uid)
	if err != nil {
		return err
	}
	if info != nil {
		return fmt.Errorf("user already exists, uid: %d", userCommand.Uid)
	}
	sensitiveInfo := entity.NewSensitiveInfo(entity.WithAge(userCommand.Age),
		entity.WithIdCard(userCommand.IdCard),
		entity.WithPassWord(userCommand.PassWord),
		entity.WithSex(userCommand.Sex),
		entity.WithPhoneNumber(userCommand.PhoneNumber))

	_, err = s.userRepository.Create(entity.NewUserInfo(userCommand.Uid,
		entity.WithAvatar(userCommand.Avatar),
		entity.WithNickName(userCommand.NickName),
		entity.WithSensitiveInfo(*sensitiveInfo)))

	return err
}

func (s *UserService) GetUserInfo(userId int64) (*entity.UserInfo, error) {
	return s.userRepository.FindByID(userId)
}

func (s *UserService) GetUserInfoByPhoneNumber(number string) (*entity.UserInfo, error) {
	return s.userRepository.FindByPhoneNumber(number)
}

func (s *UserService) UserLogin(uid int64, password string) (bool, error) {
	info, err := s.userRepository.FindByID(uid)
	if err != nil {
		return false, err
	}
	return info.SensitiveInfo.CheckUnEncryptPassWord(password)
}
