package model

type UserInfo struct {
	NickName string `json:"nick_name" gorm:"column:nick_name"`
	Uid      string `json:"uid" gorm:"column:uid;primary_key;"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	//SensitiveInfo *SensitiveInfo `json:"sensitive_info" gorm:"column:sensitive_info"`
}

type SensitiveInfo struct {
	Sex         string `json:"sex" gorm:"column:sex"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	IdCard      string `json:"id_card" gorm:"column:id_card"`
	Age         int64  `json:"age" gorm:"column:age"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
