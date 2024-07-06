package school_service

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/application/command"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
	"github.com/onebillion-0/user_sdk/biz/domain/repositories"
)

type RegisterService struct {
	Member repositories.MemberRepository
	System repositories.SystemRepository
}

func NewRegisterService(member repositories.MemberRepository, system repositories.SystemRepository) *RegisterService {
	return &RegisterService{
		Member: member,
		System: system,
	}
}

func (r *RegisterService) RegisterMembers(ctx context.Context, cmds []*command.SchoolMemberCommand) error {
	for _, memberCommand := range cmds {
		_, err := r.System.FindByAppID(ctx, memberCommand.Appid)
		if err != nil {
			return err
		}
	}
	users := make([]*school_members.Member, 0, len(cmds))
	for _, cmd := range cmds {
		user := &school_members.Member{
			NickName: cmd.NickName,
			Uid:      cmd.Uid,
			Age:      cmd.Age,
			Password: cmd.Password,
			AppId:    cmd.Appid,
			Gender:   cmd.Gender,
			Role:     cmd.Role,
		}
		users = append(users, user)
	}
	return r.Member.BatchCreate(ctx, users)
}

func (r *RegisterService) RegisterAppId() error {
	return nil
}

func (r *RegisterService) GetAllAppID(ctx context.Context) ([]int64, error) {
	models, err := r.System.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]int64, 0, len(models))
	for _, m := range models {
		result = append(result, m.AppId)
	}
	return result, nil
}
