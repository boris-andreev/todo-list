package services

import (
	"fmt"

	"todo-list/internal/configuration"
	"todo-list/internal/domain"
	"todo-list/internal/repositories"
)

type TodoService struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoService() *TodoService {
	config := configuration.GetConfig()
	todoRepository, err := repositories.NewTodoRepository(config.DbConnectionString)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (r *TodoService) AddTask(task *domain.Task) error {
	return r.todoRepository.AddTask(task)
}

func (r *TodoService) EditTask(task *domain.Task) error {
	return r.todoRepository.EditTask(task)
}

func (r *TodoService) DeleteTask(id int) error {
	return r.todoRepository.DeleteTask(id)
}

func (r *TodoService) ChangeStatus(id int, status domain.Status) error {
	return r.todoRepository.ChangeStatus(id, status)
}
