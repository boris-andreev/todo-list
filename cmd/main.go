package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"todo-list/api"
	"todo-list/internal/repository"
	"todo-list/internal/service"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../.env")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	todoService := service.New(repository.New(ctx, &wg), ctx, &wg)
	todoApi := api.NewTodoApi(ctx, &wg, todoService)

	todoApi.Start()

	/*
		_, err := todoService.Login("Ivan", "Ivan")
		if err != nil {
			log.Fatalln("wrong creds")
		}

		err = todoService.AddTask(&model.Task{
			Name: "task # 1",
			Description: "task # 1 description",
			Status: model.NotStarted,
		})
		if err != nil {
			log.Fatalln("not added")
		}

		tasks, err := todoService.GetAllTasks()
		if err != nil {
			log.Fatalln("did not get tasks")
		}

		fmt.Println(tasks)

		task := tasks[0]
		task.Description = "changed"

		err = todoService.EditTask(task)
		if err != nil {
			log.Fatalln("did not edit task")
		}

		err = todoService.ChangeStatus(task.Id, model.Closed)
		if err != nil {
			log.Fatalln("did not change status of a task ")
		}

		err = todoService.DeleteTask(task.Id)
		if err != nil {
			log.Fatalln("did not edit task")
		}

		log.Println("done")
	*/

	wg.Wait()
	fmt.Println("\nGracefull shutdown is ok")
}
