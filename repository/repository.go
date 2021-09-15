package repository

import (
	"context"
	"errors"
	"questionsandanswers/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error)
	DeleteQuestion(ctx context.Context, id string) error
	FindQuestionById(ctx context.Context, id string) (*domain.Question, error)
	FindAllQuestions(ctx context.Context) (*[]domain.Question, error)
	FindQuestionByAuthor(ctx context.Context, username string) (*[]domain.Question, error)
	UpdateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error)
}

var dbclient *mongo.Client
var collection *mongo.Collection

type MongoDbRepository struct {
}

func init() {
	dbclient, _ = GetMontoDbClient()
	collection = dbclient.Database("questionsandanswers").Collection("questions")
}

func (r MongoDbRepository) CreateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	cursor, err := collection.InsertOne(ctx, question)

	if err != nil {
		return nil, err
	}

	question.ID = cursor.InsertedID.(primitive.ObjectID)

	return question, nil
}

func (r MongoDbRepository) DeleteQuestion(ctx context.Context, id string) error {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	var result *mongo.DeleteResult
	result, err = collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrorResourceNotFound
	}

	return nil
}

func (r MongoDbRepository) FindQuestionById(ctx context.Context, id string) (*domain.Question, error) {
	var question domain.Question

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&question)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, domain.ErrorResourceNotFound
	}

	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (r MongoDbRepository) FindAllQuestions(ctx context.Context) (*[]domain.Question, error) {
	var questions []domain.Question

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}

	return &questions, nil
}

func (r MongoDbRepository) FindQuestionByAuthor(ctx context.Context, username string) (*[]domain.Question, error) {
	var questions []domain.Question

	cursor, err := collection.Find(ctx, bson.M{"author.username": username})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}

	return &questions, nil
}

func (r MongoDbRepository) UpdateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	_, err := collection.ReplaceOne(ctx, bson.M{"_id": question.ID}, question)

	if err != nil {
		return nil, err
	}

	return question, nil
}
