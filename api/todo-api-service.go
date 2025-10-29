package api

import (
	"todo-list/internal/model"
)

type todoApiService interface {
	AddTask(task *model.Task, userId int32) error
	EditTask(task *model.Task, userId int32) error
	DeleteTask(id string, userId int32) error
	ChangeStatus(id string, status model.Status, userId int32) error
	GetTaskById(id string, userId int32) (*model.Task, error)
	GetTasks(filter *model.Filter, userId int32) ([]*model.Task, error)
	GetAllTasks(userId int32) ([]*model.Task, error)
	Login(username string, password string) (userId int32, err error)
}
