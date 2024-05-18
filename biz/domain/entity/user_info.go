package entity

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserInfo struct {
	Id            int64         `json:"id"`
	NickName      string        `json:"name"`
	Avatar        string        `json:"avatar"`
	SensitiveInfo SensitiveInfo `json:"sensitive_info"`
	CreateAt      time.Time     `json:"create_at"`
	UpdateAt      time.Time     `json:"update_at"`
}

type UserOptionFunc func(*UserInfo)

func WithNickName(nickName string) UserOptionFunc {
	return func(info *UserInfo) {
		info.NickName = nickName
	}
}

func WithAvatar(avatar string) UserOptionFunc {
	return func(info *UserInfo) {
		info.Avatar = avatar
	}
}

func WithSensitiveInfo(sensitiveInfo SensitiveInfo) UserOptionFunc {
	return func(info *UserInfo) {
		info.SensitiveInfo = sensitiveInfo
	}
}

func NewUserInfo(uid int64, optionFunc ...UserOptionFunc) *UserInfo {
	userInfo := &UserInfo{Id: uid, CreateAt: time.Now(), UpdateAt: time.Now()}
	for _, fn := range optionFunc {
		fn(userInfo)
	}
	return userInfo
}

func (u *UserInfo) Validate() error {
	if u.Id == 0 {
		return errors.New("uuid is required")
	}
	if u.NickName == "" {
		return errors.New("nickname is required")
	}
	return nil
}

func (u *UserInfo) UpdateSensitiveInfo(info SensitiveInfo) error {
	u.UpdateAt = time.Now()
	u.SensitiveInfo = info
	return nil
}

type SensitiveInfo struct {
	Sex         string `json:"sex"`
	PassWord    string `json:"pass_word"`
	PhoneNumber string `json:"phone_number"`
	IdCard      string `json:"id_card"`
	Age         int    `json:"age"`
}
type SensitiveInfoOption func(*SensitiveInfo)

func WithAge(age int) SensitiveInfoOption {
	return func(info *SensitiveInfo) {
		info.Age = age
	}
}

func WithSex(sex string) SensitiveInfoOption {
	return func(info *SensitiveInfo) {
		info.Sex = sex
	}
}

func WithPhoneNumber(phoneNumber string) SensitiveInfoOption {
	return func(info *SensitiveInfo) {
		info.PhoneNumber = phoneNumber
	}
}

func WithIdCard(idCard string) SensitiveInfoOption {
	return func(info *SensitiveInfo) {
		info.IdCard = idCard
	}
}

func WithPassWord(passWord string) SensitiveInfoOption {
	return func(info *SensitiveInfo) {
		info.PassWord = passWord
	}
}

func NewSensitiveInfo(opts ...SensitiveInfoOption) *SensitiveInfo {
	info := SensitiveInfo{}
	for _, opt := range opts {
		opt(&info)
	}
	return &info
}

func (s *SensitiveInfo) Validate() error {
	if s.PassWord == "" {
		return errors.New("password is required")
	}
	return nil
}

func (s *SensitiveInfo) Encrypt() error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(s.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("encrypted error: %s", err.Error())
	}
	s.PassWord = string(encrypted)
	return nil
}

func (s *SensitiveInfo) CheckUnEncryptPassWord(unEncryptPassword string) (bool, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(s.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("encrypted error: %s", err.Error())
	}
	return string(encrypted) == s.PassWord, nil
}
