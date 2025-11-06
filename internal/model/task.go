package model

type Status int8

const (
	NotStarted Status = 1 << iota
	InProgress
	Closed
	Finished
	All = NotStarted | InProgress | Closed | Finished
)

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      Status `json:"status" binding:"required"`
}
