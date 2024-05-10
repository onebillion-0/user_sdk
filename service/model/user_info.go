package model

type CreatUserRequest struct {
	NickName      string
	Uid           string
	Avatar        string
	SensitiveInfo *SensitiveInfo
}

type SensitiveInfo struct {
	Sex         string `json:"sex" gorm:"column:sex"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	IdCard      string `json:"id_card" gorm:"column:id_card"`
	Age         int64  `json:"age" gorm:"column:age"`
}
