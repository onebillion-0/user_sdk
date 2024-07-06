package school_members

type System struct {
	Id         int64 `json:"id" bson:"id"`
	AppId      int64 `json:"app_id" bson:"app_id"`
	CreateTime int64 `json:"create_time" bson:"create_time"`
	UpdateTime int64 `json:"update_time" bson:"update_time"`
}
