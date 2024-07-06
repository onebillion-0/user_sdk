package mongodb

import (
	"context"
	"fmt"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
	"github.com/onebillion-0/user_sdk/biz/domain/repositories"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoSystemRepository struct {
	collection *mongo.Collection
	flake      *sonyflake.Sonyflake
}

func NewMongoSystemRepository(db *mongo.Database, collectionName string) repositories.SystemRepository {
	// 初始化sonyflake
	st := sonyflake.Settings{}
	flake := sonyflake.NewSonyflake(st)
	if flake == nil {
		panic(fmt.Errorf("sonyflake not created"))
	}
	return &MongoSystemRepository{
		collection: db.Collection(collectionName),
		flake:      flake,
	}
}

func (repo *MongoSystemRepository) FindByAppID(ctx context.Context, appid int64) (*school_members.System, error) {
	var system school_members.System
	filter := bson.D{{"app_id", appid}}
	err := repo.collection.FindOne(ctx, filter).Decode(&system)
	if err != nil {
		return nil, err
	}
	return &system, nil
}

func (repo *MongoSystemRepository) CreateByAppID(ctx context.Context, appid int64) (*school_members.System, error) {

	now := time.Now().Unix()
	newSystem := &school_members.System{
		AppId:      appid,
		CreateTime: now,
		UpdateTime: now,
	}
	_, err := repo.collection.InsertOne(ctx, newSystem)
	if err != nil {
		return nil, err
	}
	return newSystem, nil
}

func (repo *MongoSystemRepository) Update(ctx context.Context, sys *school_members.System) (*school_members.System, error) {
	sys.UpdateTime = time.Now().Unix()
	filter := bson.D{{"app_id", sys.AppId}}
	update := bson.D{{"$set", sys}}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return sys, nil
}

func (repo *MongoSystemRepository) Delete(ctx context.Context, appid int64) error {
	filter := bson.D{{"app_id", appid}}
	_, err := repo.collection.DeleteOne(ctx, filter)
	return err
}
