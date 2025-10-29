package model

type Repository interface {
	AddTask(task *Task, userId int32) error
	EditTask(task *Task, userId int32) error
	DeleteTask(id string, userId int32) error
	ChangeStatus(id string, status Status, userId int32) error
	GetTaskById(id string, userId int32) (*Task, error)
	GetTasks(filter Filter, userId int32) ([]*Task, error)
	GetAllTasks(userId int32) ([]*Task, error)
	GetUserByName(username string) (*User, error)
}
