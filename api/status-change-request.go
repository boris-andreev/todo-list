package api

import "todo-list/internal/model"

type statusChangeRequest struct {
	Id     string
	Status model.Status
}
