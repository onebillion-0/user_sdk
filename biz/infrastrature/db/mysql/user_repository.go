package mysql

import (
	"errors"
	"fmt"
	"github.com/oneliuliu/user_sdk/biz/domain/entity"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) FindByID(uid int64) (*entity.UserInfo, error) {
	var info UserInfo
	if err := repo.db.Preload("user_info").Where("id = ?", uid).First(&info).Error; err != nil {
		return nil, err
	}
	return fromDBUserInfo(&info), nil
}

func (repo *GormUserRepository) FindByPhoneNumber(number string) (*entity.UserInfo, error) {
	var info UserInfo
	if err := repo.db.Preload("user_info").Where("phone_number = ?", number).First(&info).Error; err != nil {
		return nil, err
	}
	return fromDBUserInfo(&info), nil
}

func (repo *GormUserRepository) Create(user *entity.UserInfo) (*entity.UserInfo, error) {
	info := toDBUserInfo(user)
	err := repo.db.AutoMigrate(&UserInfo{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := repo.db.Create(&info).Error; err != nil {
		return nil, err
	}
	return repo.FindByID(info.Uid)
}

func (repo *GormUserRepository) Update(user *entity.UserInfo) (*entity.UserInfo, error) {
	info := toDBUserInfo(user)
	if err := repo.db.Model(&UserInfo{}).Where("id = ?", info.Uid).Updates(info).Error; err != nil {
		return nil, err
	}
	return repo.FindByID(info.Uid)
}

func (repo *GormUserRepository) Delete(id int64) error {
	return errors.New("delete user info is not supported")
}
