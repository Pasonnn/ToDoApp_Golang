package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todo-api/models"
	"todo-api/services"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	taskService *services.TaskService
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// CreateTask handles POST /tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
		})
		return
	}

	task, err := h.taskService.CreateTask(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create task",
			Data:       nil,
		})
		return
	}

	response := models.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Task created successfully",
		Data:       response,
	})
}

// GetAllTasks handles GET /tasks
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks := h.taskService.GetAllTasks()

	// Convert to response format
	response := make([]models.TaskResponse, len(tasks))
	for i, task := range tasks {
		response[i] = models.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		}
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Tasks retrieved successfully",
		Data:       response,
	})
}

// UpdateTask handles PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid task ID",
			Data:       nil,
		})
		return
	}

	var req models.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
		})
		return
	}

	task, err := h.taskService.UpdateTask(id, &req)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success:    false,
				StatusCode: http.StatusNotFound,
				Message:    "Task not found",
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update task",
			Data:       nil,
		})
		return
	}

	response := models.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Task updated successfully",
		Data:       response,
	})
}

// DeleteTask handles DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid task ID",
			Data:       nil,
		})
		return
	}

	err = h.taskService.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success:    false,
				StatusCode: http.StatusNotFound,
				Message:    "Task not found",
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete task",
			Data:       nil,
		})
		return
	}

	response := models.DeleteTaskResponse{
		ID: id,
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Task deleted successfully",
		Data:       response,
	})
}
