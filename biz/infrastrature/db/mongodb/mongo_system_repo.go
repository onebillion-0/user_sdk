package mongodb

import (
	"context"
	"fmt"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
	"github.com/onebillion-0/user_sdk/biz/domain/repositories"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type MongoSystemRepository struct {
	collection *mongo.Collection
	flake      *sonyflake.Sonyflake
}

var once sync.Once

func NewMongoSystemRepository(db *mongo.Database, collectionName string) repositories.SystemRepository {
	// 初始化sonyflake
	st := sonyflake.Settings{}
	flake := sonyflake.NewSonyflake(st)
	if flake == nil {
		panic(fmt.Errorf("sonyflake not created"))
	}
	repo := &MongoSystemRepository{
		collection: db.Collection(collectionName),
		flake:      flake,
	}
	once.Do(repo.init)
	return repo
}

func (repo *MongoSystemRepository) init() {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := repo.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
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

func (repo *MongoSystemRepository) CreateByAppID(ctx context.Context, appid int64, name string) (*school_members.System, error) {

	now := time.Now().Unix()
	newSystem := &school_members.System{
		AppId:      appid,
		SystemName: name,
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

func (repo *MongoSystemRepository) GetAll(ctx context.Context) ([]*school_members.System, error) {
	cusor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cusor.Close(ctx)
	result := make([]*school_members.System, 0)
	err = cusor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
