package user_sdk

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/application/command"
	"github.com/onebillion-0/user_sdk/biz/application/services/school_service"
	"github.com/onebillion-0/user_sdk/biz/infrastrature/db/mongodb"
	"github.com/onebillion-0/user_sdk/biz/infrastrature/db/mongodb/mongo_table"
	"github.com/onebillion-0/user_sdk/biz/interface/sdk"
)

func SchoolMemberLogin(ctx context.Context, uid int64, password string) (token string, err error) {
	mongoRepo := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sev := school_service.NewLoginService(mongoRepo)
	return sdk.NewSchoolLoginController(sev).Login(ctx, uid, password)
}

func SchoolMemberRegister(ctx context.Context, command []*command.SchoolMemberCommand) error {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).RegisterMembers(ctx, command)
}

func SchoolSystemRegister(ctx context.Context, appid int64) error {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).RegisterAppId(ctx, appid)
}

func GetAppIDList(ctx context.Context) (appIDList []int64, err error) {
	member := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName())
	sys := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	return school_service.NewRegisterService(member, sys).GetAllAppID(ctx)
}
