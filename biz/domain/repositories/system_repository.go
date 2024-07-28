package repositories

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
)

type SystemRepository interface {
	FindByAppID(ctx context.Context, id int64) (*school_members.System, error)
	CreateByAppID(ctx context.Context, appid int64, name string) (*school_members.System, error)
	Update(ctx context.Context, sys *school_members.System) (*school_members.System, error)
	Delete(ctx context.Context, appid int64) error
	GetAll(ctx context.Context) ([]*school_members.System, error)
}

type ClassRepository interface {
	GetAll(ctx context.Context) ([]*school_members.Class, error)
}
