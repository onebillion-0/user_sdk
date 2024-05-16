package postgres

import (
	"github.com/oneliuliu/user_sdk/biz/domain/entity"
)

func fromDBUserInfo(info *UserInfo) *entity.UserInfo {
	return &entity.UserInfo{
		Id:            info.Uid,
		NickName:      info.NickName,
		Avatar:        info.Avatar,
		CreateAt:      info.CreatedAt,
		UpdateAt:      info.UpdatedAt,
		SensitiveInfo: *fromDBSensitive(&info.SensitiveInfo),
	}
}

func fromDBSensitive(s *SensitiveInfo) *entity.SensitiveInfo {
	return &entity.SensitiveInfo{
		Sex:         s.Sex,
		PassWord:    s.PassWord,
		PhoneNumber: s.PhoneNumber,
		IdCard:      s.IdCard,
		Age:         s.Age,
	}
}

func toDBUserInfo(in *entity.UserInfo) *UserInfo {
	return &UserInfo{
		Uid:           in.Id,
		NickName:      in.NickName,
		Avatar:        in.Avatar,
		CreatedAt:     in.CreateAt,
		UpdatedAt:     in.UpdateAt,
		SensitiveInfo: *toDBSensitive(&in.SensitiveInfo),
	}
}

func toDBSensitive(in *entity.SensitiveInfo) *SensitiveInfo {
	return &SensitiveInfo{
		Sex:         in.Sex,
		PassWord:    in.PassWord,
		PhoneNumber: in.PhoneNumber,
		IdCard:      in.IdCard,
		Age:         in.Age,
	}
}
