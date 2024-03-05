package repository

import (
	"ApiGo/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks//repository/mockTodoRepository.go -package=repository ApiGo/Repository TodoRepostory
type TodoRepostoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepostory interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([](models.Todo), error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepostoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		errors.New("Failed to insert todo")
		return false, err
	}
	return true, nil
}

func (t TodoRepostoryDB) GetAll() ([](models.Todo), error) {
	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func (t TodoRepostoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDb(dbClient *mongo.Collection) TodoRepostoryDB {
	return TodoRepostoryDB{TodoCollection: dbClient}
}
