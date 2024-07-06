package school_members

type SuperAdmin struct {
	Id         int64  `json:"id" bson:"id"`
	Uid        int64  `json:"uid" bson:"uid"`
	Name       string `json:"name" bson:"name"`
	Age        int64  `json:"age" bson:"age"`
	Password   string `json:"password" bson:"password"`
	Gender     string `json:"gender" bson:"gender"`
	AppId      int64  `json:"app_id" bson:"app_id"`
	CreateTime int64  `json:"create_time" bson:"create_time"`
	UpdateTime int64  `json:"update_time" bson:"update_time"`
}
