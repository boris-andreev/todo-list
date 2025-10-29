package api

import (
	"net/http"
	"todo-list/internal/model"
	"todo-list/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoApiHandler struct {
	todoApiService todoApiService
}

func NewTodoApiHandler(todoApiService todoApiService) *TodoApiHandler {
	return &TodoApiHandler{
		todoApiService: todoApiService,
	}
}

func (h *TodoApiHandler) AddTask() gin.HandlerFunc {
	return func(c *gin.Context) {

		task := &model.Task{}
		h.createOrUpdate(c, task, h.todoApiService.AddTask)
	}
}

func (h *TodoApiHandler) createOrUpdate(c *gin.Context, task *model.Task, mutator func(task *model.Task, userId int32) error) {
	userId, ok := parseUserId(c)
	if !ok {
		return
	}
	err := c.ShouldBindJSON(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{
			Message: err.Error(),
		})

		return
	}

	err = mutator(task, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage{
			Message: err.Error(),
		})
	}
}

func parseUserId(c *gin.Context) (int32, bool) {
	s, ok := c.Get(service.UserIdKey)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return 0, false
	}

	userId, ok := s.(int32)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return 0, false
	}

	return int32(userId), true
}
