package domain

import "time"

type Status int8

const (
	NotStarted Status = iota
	InProgress
	Closed
	Finished
)

type Task struct {
	Id          int
	Name        string
	Description string
	Status      Status
	CreatedAt   time.Time
}
