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

type MongoStudentRepository struct {
	collection *mongo.Collection
}

func NewMongoMemberRepository(db *mongo.Database, collectionName string) repositories.MemberRepository {
	return &MongoStudentRepository{
		collection: db.Collection(collectionName),
	}
}

func (repo *MongoStudentRepository) FindByID(ctx context.Context, uid int64) (*school_members.Member, error) {
	var student school_members.Member
	filter := bson.D{{"uid", uid}}
	err := repo.collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (repo *MongoStudentRepository) FindByUsername(ctx context.Context, username string) (*school_members.Member, error) {
	var student school_members.Member
	filter := bson.D{{"name", username}}
	err := repo.collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (repo *MongoStudentRepository) Create(ctx context.Context, student *school_members.Member) (*school_members.Member, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	student.Password = string(hashedPassword)
	student.CreateTime = time.Now().Unix()
	student.UpdateTime = time.Now().Unix()

	_, err = repo.collection.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (repo *MongoStudentRepository) Update(ctx context.Context, student *school_members.Member) (*school_members.Member, error) {
	student.UpdateTime = time.Now().Unix()
	filter := bson.D{{"uid", student.Uid}}
	update := bson.D{{"$set", student}}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return student, nil
}
