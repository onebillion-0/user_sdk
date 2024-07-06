package repositories

import (
	"context"
	"github.com/oneliuliu/user_sdk/biz/domain/entity/school_members"
)

type AdminRepository interface {
	FindByID(ctx context.Context, uid int64) (*school_members.SuperAdmin, error)
	FindByUsername(ctx context.Context, username string) (*school_members.SuperAdmin, error)
	Create(ctx context.Context, user *school_members.SuperAdmin) (*school_members.SuperAdmin, error)
	Update(ctx context.Context, user *school_members.SuperAdmin) (*school_members.SuperAdmin, error)
}
