package school_members

type System struct {
	AppId      int64 `json:"app_id"        bson:"app_id"`
	SystemName int64 `json:"system_name"   bson:"system_name"`
	CreateTime int64 `json:"create_time"   bson:"create_time"`
	UpdateTime int64 `json:"update_time"   bson:"update_time"`
}

type Class struct {
	Name       string `json:"name"        bson:"name"`
	ClassId    int64  `json:"class_id"    bson:"class_id"`
	SystemId   int64  `json:"system_id"   bson:"system_id"`
	CreateTime int64  `json:"create_time" bson:"create_time"`
	UpdateTime int64  `json:"update_time" bson:"update_time"`
}
