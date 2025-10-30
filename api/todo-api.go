package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"todo-list/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoApi struct {
	router *gin.Engine
	server *http.Server
	ctx    context.Context
	wg     *sync.WaitGroup
}

func configureRouting(router *gin.Engine, handler *TodoApiHandler) {

	router.POST("/login", handler.Login())

	router.Use(HandleAuth())

	apiGroup := router.Group("/api")

	taskGroup := apiGroup.Group("/task")
	taskGroup.POST("/", handler.AddTask())
	taskGroup.PUT("/", handler.EditTask())
	taskGroup.DELETE("/:id", handler.DeleteTask())
	taskGroup.GET("/:id", handler.GetTaskById())
	taskGroup.GET("/", handler.GetAllTasks())
	taskGroup.GET("/filter", handler.GetTasks())
	taskGroup.POST("/status", handler.ChangeStatus())
}

func NewTodoApi(ctx context.Context, wg *sync.WaitGroup, todoService *service.TodoService) *TodoApi {
	router := gin.Default()
	todoApiHandler := NewTodoApiHandler(todoService)
	configureRouting(router, todoApiHandler)

	return &TodoApi{
		router: router,
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		ctx: ctx,
		wg:  wg,
	}
}

func (a *TodoApi) Start() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("Server failed to start: %v", err)
		}
	}()

	a.listenForFinish()
}

func (a *TodoApi) listenForFinish() {
	a.wg.Add(1)

	go func() {
		defer a.wg.Done()

		for {
			select {
			case <-a.ctx.Done():
				log.Println("Shutting down server...")

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := a.server.Shutdown(ctx); err != nil {
					log.Printf("Server forced to shutdown: %v", err)
				}

				log.Println("Server exited")
				return
			}
		}
	}()

}
