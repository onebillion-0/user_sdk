package mongodb

import (
	"context"
	"github.com/oneliuliu/user_sdk/biz/domain/entity/school_members"
	"github.com/oneliuliu/user_sdk/biz/domain/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type MongoAdminRepository struct {
	collection *mongo.Collection
}

func NewMongoAdminRepository(db *mongo.Database, collectionName string) repositories.AdminRepository {
	return &MongoAdminRepository{
		collection: db.Collection(collectionName),
	}
}

func (repo *MongoAdminRepository) FindByID(ctx context.Context, uid int64) (*school_members.SuperAdmin, error) {
	var admin school_members.SuperAdmin
	filter := bson.D{{"uid", uid}}
	err := repo.collection.FindOne(context.TODO(), filter).Decode(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repo *MongoAdminRepository) Create(ctx context.Context, user *school_members.SuperAdmin) (*school_members.SuperAdmin, error) {
	user.CreateTime = time.Now().Unix()
	user.UpdateTime = time.Now().Unix()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	_, err = repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *MongoAdminRepository) FindByUsername(ctx context.Context, username string) (*school_members.SuperAdmin, error) {
	var admin school_members.SuperAdmin
	filter := bson.D{{"name", username}}
	err := repo.collection.FindOne(context.TODO(), filter).Decode(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repo *MongoAdminRepository) Update(ctx context.Context, user *school_members.SuperAdmin) (*school_members.SuperAdmin, error) {
	user.UpdateTime = time.Now().Unix()
	filter := bson.D{{"uid", user.Uid}}
	update := bson.D{{"$set", user}}
	_, err := repo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}
