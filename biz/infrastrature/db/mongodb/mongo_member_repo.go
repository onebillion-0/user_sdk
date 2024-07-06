package mongodb

import (
	"context"
	"github.com/onebillion-0/user_sdk/biz/domain/entity/school_members"
	"github.com/onebillion-0/user_sdk/biz/domain/repositories"
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

func (repo *MongoStudentRepository) BatchCreate(ctx context.Context, users []*school_members.Member) error {
	users, err := repo.initMemberModels(users)
	if err != nil {
		return err
	}

	models := make([]mongo.WriteModel, 0, len(users))
	for _, user := range users {
		model := mongo.NewInsertOneModel()
		model.SetDocument(user)
		models = append(models, model)
	}
	_, err = repo.collection.BulkWrite(ctx, models)
	return err
}

func (repo *MongoStudentRepository) initMemberModels(users []*school_members.Member) ([]*school_members.Member, error) {
	now := time.Now().Unix()
	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
		user.CreateTime = now
		user.UpdateTime = now
	}
	return users, nil
}
