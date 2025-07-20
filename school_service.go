package user_sdk

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/application/command"
	"github.com/onebillion-0/user_sdk/biz/application/services/school_service"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
	"github.com/onebillion-0/user_sdk/biz/infrastrature/db/mongodb"
	"github.com/onebillion-0/user_sdk/biz/infrastrature/db/mongodb/mongo_table"
	"github.com/onebillion-0/user_sdk/biz/interface/sdk"
)

func SchoolMemberLogin(ctx context.Context, uid, appid int64, password string) (token string, cmd *command.SchoolMemberCommand, err error) {
	mongoRepo := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sev := school_service.NewLoginService(mongoRepo)
	return sdk.NewSchoolLoginController(sev).Login(ctx, uid, appid, password)
}

func SchoolMemberRegister(ctx context.Context, command []*command.SchoolMemberCommand) error {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).RegisterMembers(ctx, command)
}

func MGetStudents(ctx context.Context, page int, size int) (int, []command.SchoolMemberCommand, error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).MGetStudents(ctx, page, size)
}

func SchoolSystemRegister(ctx context.Context, appid int64, name string) error {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).RegisterAppId(ctx, appid, name)
}

func GetAppIDList(ctx context.Context) (appIDList map[int64]string, err error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).GetAllAppID(ctx)
}

func GetRoleByID(ctx context.Context, id int64) (school_members.Role, error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).GetRoleById(ctx, id)
}

func GetUserInfoByID(ctx context.Context, id int64) (*school_members.Member, error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).GetUserInfoByID(ctx, id)
}

func BatchGetUserInfo(ctx context.Context, ids []int64) ([]*school_members.Member, error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).BatchGetUser(ctx, ids)
}

// 1111
func DeleteMember(ctx context.Context, ids []int64) error {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).DeleteMember(ctx, ids)
}

func CheckExpireTime(ctx context.Context, uid, aid int64) (bool, error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).CheckExpireTime(ctx, uid, aid)
}
