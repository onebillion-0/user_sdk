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

func NewUserInfo() *UserInfo {
	return &UserInfo{}
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
	Age         int64  `json:"age"`
}

func NewSensitiveInfo() *SensitiveInfo {
	return &SensitiveInfo{}
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
