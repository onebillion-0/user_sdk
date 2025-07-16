package command

import "github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"

type SchoolMemberCommand struct {
	NickName   string
	Uid        int64
	Age        int64
	ClassId    int64
	Password   string
	Appid      int64
	Gender     string
	Role       school_members.Role
	ExpireTime int64 //秒级别时间戳
}
