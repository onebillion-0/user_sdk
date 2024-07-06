package repositories

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
)

type MemberRepository interface {
	FindByID(ctx context.Context, uid int64) (*school_members.Member, error)
	FindByUsername(ctx context.Context, name string) (*school_members.Member, error)
	Create(ctx context.Context, user *school_members.Member) (*school_members.Member, error)
	Update(ctx context.Context, user *school_members.Member) (*school_members.Member, error)
}
