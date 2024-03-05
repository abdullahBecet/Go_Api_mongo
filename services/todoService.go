package services

import (
	"ApiGo/dto"
	"ApiGo/models"
	"ApiGo/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks//service/mockTodoservice.go -package=services ApiGo/services TodoService
type DefaultTodoServices struct {
	Repo repository.TodoRepostory
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoServices) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}
	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.TodoDTO{Status: result}
	return &res, nil
}
func (t DefaultTodoServices) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (t DefaultTodoServices) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)
	if err != nil || result == false {
		return false, err
	}
	return true, nil
}

func NewTodoService(Repo repository.TodoRepostory) DefaultTodoServices {
	return DefaultTodoServices{Repo: Repo}
}
