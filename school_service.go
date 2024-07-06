package user_sdk

import (
	"context"
	"github.com/oneliuliu/user_sdk/biz/application/services/school_service"
	"github.com/oneliuliu/user_sdk/biz/infrastrature/db/mongodb"
	"github.com/oneliuliu/user_sdk/biz/infrastrature/db/mongodb/mongo_table"
	"github.com/oneliuliu/user_sdk/biz/interface/sdk"
)

func SchoolMemberLogin(ctx context.Context, appid int64, uid int64, password string) (token string, err error) {
	mongoRepo := mongodb.NewMongoMemberRepository(mongodb.MongoClient, mongo_table.GetMemberCollectionName(appid))
	sysMongoRepo := mongodb.NewMongoSystemRepository(mongodb.MongoClient, mongo_table.GetSysCollectionName())
	sev := school_service.NewLoginService(mongoRepo, sysMongoRepo)
	return sdk.NewSchoolLoginController(sev).Login(ctx, uid, password)
}
