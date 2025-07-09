package repositories

import (
	"errors"

	"todo-list/internal/domain"
)

type TodoRepository struct {
	dbConnectionString string
}

func (r *TodoRepository) AddTask(task *domain.Task) error {
	return errors.New("Not implemented")
}

func (r *TodoRepository) EditTask(task *domain.Task) error {
	return errors.New("Not implemented")
}

func (r *TodoRepository) DeleteTask(id int) error {
	return errors.New("Not implemented")
}

func (r *TodoRepository) ChangeStatus(id int, status domain.Status) error {
	return errors.New("Not implemented")
}

func (r *TodoRepository) GetTaskById(id int) (*domain.Task, error) {
	return nil, errors.New("Not implemented")
}

func (r *TodoRepository) GetTasks() ([]domain.Task, error) {
	return nil, errors.New("Not implemented")
}

func NewTodoRepository(dbConnectionString string) (*TodoRepository, error) {
	return &TodoRepository{
		dbConnectionString: dbConnectionString,
	}, nil
}
