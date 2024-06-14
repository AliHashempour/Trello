package handler

import (
	"Trello/internal/model"
	"Trello/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SubTaskHandler struct {
	repo repository.SubTaskRepository
}

func NewSubTaskHandler(repo repository.SubTaskRepository) *SubTaskHandler {
	return &SubTaskHandler{repo: repo}
}

func (h *SubTaskHandler) Register(g *echo.Group) {
	g.GET("/", h.GetAllSubTasks)
	g.POST("/", h.CreateSubTask)
	g.GET("/:subTaskId", h.GetSubTask)
	g.PUT("/:subTaskId", h.UpdateSubTask)
	g.DELETE("/:subTaskId", h.DeleteSubTask)
}

func (h *SubTaskHandler) GetAllSubTasks(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid task ID")
	}
	subTasks, err := h.repo.GetAll(uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, subTasks)
}

func (h *SubTaskHandler) GetSubTask(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid task ID")
	}
	subTaskID, err := strconv.ParseUint(c.Param("subTaskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid subtask ID")
	}
	subTask, err := h.repo.GetBy(map[string]interface{}{"id": subTaskID, "task_id": taskID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, subTask)
}

func (h *SubTaskHandler) CreateSubTask(c echo.Context) error {
	var subTask model.SubTask
	if err := c.Bind(&subTask); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	taskId, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task ID format")
	}

	subTask.TaskID = uint(taskId)

	if err := h.repo.Create(&subTask); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, subTask)

}

func (h *SubTaskHandler) UpdateSubTask(c echo.Context) error {
	subTaskID, err := strconv.ParseUint(c.Param("subTaskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid subtask ID")
	}
	var subTask model.SubTask
	if err := c.Bind(&subTask); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	subTask.ID = uint(subTaskID)
	if err := h.repo.Update(&subTask); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, subTask)
}

func (h *SubTaskHandler) DeleteSubTask(c echo.Context) error {
	subTaskID, err := strconv.ParseUint(c.Param("subTaskId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid subtask ID")
	}
	if err := h.repo.DeleteBy(map[string]interface{}{"id": subTaskID}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Subtask deleted successfully")
}
