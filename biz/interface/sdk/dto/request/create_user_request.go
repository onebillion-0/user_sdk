package request

import (
	"fmt"
	"github.com/onebillion-0/user_sdk/biz/application/command"
	"strconv"
)

type CreateUserRequest struct {
	Uid         string `json:"uid"`
	NickName    string `json:"nick_name"`
	Avatar      string `json:"avatar"`
	Sex         string `json:"sex"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	IdCard      string `json:"id_card"`
	Age         string `json:"age"`
}

func (c CreateUserRequest) ToCreateUserCommand() (*command.CreateUserCommand, error) {
	uid, err := strconv.ParseInt(c.Uid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("uid format incorrect, uid: %v, err: %w", c.Uid, err)
	}
	var age int
	if c.Age != "" {
		//todo :age
		age, _ = strconv.Atoi(c.Age)
	}
	return &command.CreateUserCommand{
		Uid:         uid,
		NickName:    c.NickName,
		Avatar:      c.Avatar,
		Sex:         c.Sex,
		PassWord:    c.Password,
		PhoneNumber: c.PhoneNumber,
		IdCard:      c.IdCard,
		Age:         age,
	}, nil
}
