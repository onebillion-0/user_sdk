package school_members

type Role string

const Student Role = "student"
const Admin Role = "admin"
const Teacher Role = "teacher"

type Member struct {
	Uid        int64  `json:"uid" bson:"uid"`
	ClassId    int64  `json:"class_id" bson:"class_id"`
	NickName   string `json:"nick_name" bson:"nick_name"`
	Age        int64  `json:"age" bson:"age"`
	Password   string `json:"password" bson:"password"`
	Gender     string `json:"gender" bson:"gender"`
	AppId      int64  `json:"app_id" bson:"app_id"`
	Role       Role   `json:"role" bson:"role"`
	CreateTime int64  `json:"create_time" bson:"create_time"`
	UpdateTime int64  `json:"update_time" bson:"update_time"`
}
