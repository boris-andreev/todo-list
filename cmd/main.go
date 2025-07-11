package main

import (
	"fmt"

	"todo-list/internal/domain"
	"todo-list/internal/services"
)

func main() {
	todoService := services.NewTodoService()
	err := todoService.AddTask(&domain.Task{
		Name: "Fake one",
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
