package service

import (
	"context"
	"sync"

	"todo-list/internal/model"

	"golang.org/x/crypto/bcrypt"
)

const (
	UserIdKey = "userId"
)

type TodoService struct {
	todoRepository model.Repository
	ctx            context.Context
	wg             *sync.WaitGroup
}

func (s *TodoService) AddTask(task *model.Task, userId int32) error {
	return s.todoRepository.AddTask(task, userId)
}

func (s *TodoService) EditTask(task *model.Task, userId int32) error {
	return s.todoRepository.EditTask(task, userId)
}

func (s *TodoService) DeleteTask(id string, userId int32) error {
	return s.todoRepository.DeleteTask(id, userId)
}

func (s *TodoService) ChangeStatus(id string, status model.Status, userId int32) error {
	return s.todoRepository.ChangeStatus(id, status, userId)
}

func (s *TodoService) GetTaskById(id string, userId int32) (*model.Task, error) {
	return s.todoRepository.GetTaskById(id, userId)
}

func (s *TodoService) GetTasks(filter *model.Filter, userId int32) ([]*model.Task, error) {
	return s.todoRepository.GetTasks(filter, userId)
}

func (s *TodoService) GetAllTasks(userId int32) ([]*model.Task, error) {
	return s.todoRepository.GetAllTasks(userId)
}

func (s *TodoService) Login(username string, password string) (userId int32, err error) {
	user, err := s.todoRepository.GetUserByName(username)

	if err != nil {
		return 0, err
	}

	err = checkPasswordHash(password, user.Password)

	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
func New(todoRepository model.Repository, ctx context.Context, wg *sync.WaitGroup) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
		ctx:            ctx,
		wg:             wg,
	}
}
