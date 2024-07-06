package mongo_table

import "strconv"

func GetMemberCollectionName(appid int64) string {
	return "school_member_" + strconv.FormatInt(appid, 10)
}
func GetSysCollectionName() string {
	return "school_system"
}
