package api

import (
	"net/http"
	"todo-list/internal/model"
	"todo-list/internal/service"
	"todo-list/internal/utils/jwt"

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

func (h *TodoApiHandler) EditTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		task := &model.Task{}
		h.createOrUpdate(c, task, h.todoApiService.EditTask)
	}
}

func (h *TodoApiHandler) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := getIdFromRoute(c)
		if !ok {
			return
		}

		userId, ok := getUserIdFromContext(c)
		if !ok {
			return
		}

		err := h.todoApiService.DeleteTask(id, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
			})

			return
		}
	}
}

func (h *TodoApiHandler) GetTaskById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := getIdFromRoute(c)
		if !ok {
			return
		}

		userId, ok := getUserIdFromContext(c)
		if !ok {
			return
		}

		res, err := h.todoApiService.GetTaskById(id, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func (h *TodoApiHandler) GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := getUserIdFromContext(c)
		if !ok {
			return
		}

		res, err := h.todoApiService.GetAllTasks(userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *TodoApiHandler) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		req := &loginRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		userId, err := h.todoApiService.Login(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Message: "Invalid username or password",
			})

			return
		}

		token, err := jwt.GenerateToken(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
		})
	}
}

func (h *TodoApiHandler) GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := getUserIdFromContext(c)
		if !ok {
			return
		}

		filter := &model.Filter{}

		err := c.ShouldBindJSON(filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Message: err.Error(),
			})

			return
		}

		res, err := h.todoApiService.GetTasks(filter, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *TodoApiHandler) ChangeStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := getUserIdFromContext(c)
		if !ok {
			return
		}

		req := &statusChangeRequest{}

		err := c.ShouldBindJSON(req)

		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Message: err.Error(),
			})

			return
		}

		err = h.todoApiService.ChangeStatus(req.Id, req.Status, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Message: err.Error(),
			})

			return
		}
	}
}

func (h *TodoApiHandler) createOrUpdate(c *gin.Context, task *model.Task, mutator func(task *model.Task, userId int32) error) {
	userId, ok := getUserIdFromContext(c)
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

func getUserIdFromContext(c *gin.Context) (int32, bool) {
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

func getIdFromRoute(c *gin.Context) (string, bool) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, ErrorMessage{
			Message: "Task id is a must",
		})

		return "", false
	}

	return id, true
}
