package repository

import (
	"ApiGo/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepostoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepostory interface {
	Insert(todo models.Todo) (bool, error)
}

func (t TodoRepostoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		errors.New("Failed to insert todo")
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDb(dbClient *mongo.Collection) TodoRepostoryDB {
	return TodoRepostoryDB{TodoCollection: dbClient}
}
