package handler

import (
	"Trello/internal/model"
	"Trello/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTask(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) Register(g *echo.Group) {
	g.GET("/", h.GetAllTasks)
	g.POST("/", h.CreateTask)
	g.GET("/:taskId", h.GetTask)
	g.PUT("/:taskId", h.UpdateTask)
	g.DELETE("/:taskId", h.DeleteTask)
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid workspace ID")
	}

	tasks, err := h.repo.GetAll(uint(workspaceID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid workspace ID")
	}

	taskID, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid task ID")
	}

	task, err := h.repo.GetBy(map[string]interface{}{"id": taskID, "workspace_id": workspaceID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	workspaceID, _ := strconv.ParseUint(c.Param("workspaceId"), 10, 64)

	task.WorkspaceID = uint(workspaceID)

	if err := h.repo.Create(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"task": task})
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid task ID")
	}
	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.ID = uint(taskID)
	if err := h.repo.Update(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid task ID")
	}
	if err := h.repo.DeleteBy(map[string]interface{}{"id": taskID}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Task deleted successfully")
}
