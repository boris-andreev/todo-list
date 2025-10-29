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
	Id          string
	Name        string
	Description string
	Status      Status
}
