package postgres

import "time"

type UserInfo struct {
	NickName      string        `json:"nick_name" gorm:"column:nick_name"`
	Uid           int64         `json:"uid" gorm:"column:uid;primary_key;"`
	Avatar        string        `json:"avatar" gorm:"column:avatar"`
	SensitiveInfo SensitiveInfo `json:"sensitive_info" gorm:"column:sensitive_info"`
	CreatedAt     time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"column:updated_at"`
}

type SensitiveInfo struct {
	Sex         string `json:"sex" gorm:"column:sex"`
	PassWord    string `json:"pass_word" gorm:"column:pass_word"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	IdCard      string `json:"id_card" gorm:"column:id_card"`
	Age         int64  `json:"age" gorm:"column:age"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
