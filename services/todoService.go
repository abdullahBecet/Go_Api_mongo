package services

import (
	"ApiGo/dto"
	"ApiGo/models"
	"ApiGo/repository"
)

type DefaultTodoServices struct {
	Repo repository.TodoRepostory
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
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

func NewTodoService(Repo repository.TodoRepostory) DefaultTodoServices {
	return DefaultTodoServices{Repo: Repo}
}
