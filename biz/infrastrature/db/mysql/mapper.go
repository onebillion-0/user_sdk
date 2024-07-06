package mysql

import (
	"github.com/onebillion-0/user_sdk/biz/domain/entity"
)

func fromDBUserInfo(info *UserInfo) *entity.UserInfo {
	return &entity.UserInfo{
		Id:       info.Uid,
		NickName: info.NickName,
		Avatar:   info.Avatar,
		SensitiveInfo: entity.SensitiveInfo{
			Sex:         info.Sex,
			PassWord:    info.PassWord,
			PhoneNumber: info.PhoneNumber,
			IdCard:      info.IdCard,
			Age:         info.Age,
		},
		CreateAt: info.CreatedAt,
		UpdateAt: info.UpdatedAt,
	}
}

func toDBUserInfo(in *entity.UserInfo) *UserInfo {
	return &UserInfo{
		Uid:         in.Id,
		NickName:    in.NickName,
		Avatar:      in.Avatar,
		Extra:       in.Extra,
		Sex:         in.SensitiveInfo.Sex,
		PassWord:    in.SensitiveInfo.PassWord,
		PhoneNumber: in.SensitiveInfo.PhoneNumber,
		IdCard:      in.SensitiveInfo.IdCard,
		Age:         in.SensitiveInfo.Age,
		CreatedAt:   in.CreateAt,
		UpdatedAt:   in.UpdateAt,
	}
}
