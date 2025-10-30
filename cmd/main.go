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
)

func main() {

	// uncomment in debug mode
	// godotenv.Load("../.env")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	todoService := service.New(repository.New(ctx, &wg), ctx, &wg)
	todoApi := api.NewTodoApi(ctx, &wg, todoService)

	todoApi.Start()

	wg.Wait()
	fmt.Println("\nGracefull shutdown is ok")
}
