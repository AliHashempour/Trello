package handler

import (
	"net/http"
	"strconv"

	"Trello/internal/model"
	"Trello/internal/repository"
	"github.com/labstack/echo/v4"
)

type UserWorkspaceRoleHandler struct {
	repo repository.UserWorkspace
}

func NewUserWorkspaceRoleHandler(repo repository.UserWorkspace) *UserWorkspaceRoleHandler {
	return &UserWorkspaceRoleHandler{repo: repo}
}

func (h *UserWorkspaceRoleHandler) Register(g *echo.Group) {
	g.GET("/", h.GetUsersByWorkspaceID)
	g.POST("/", h.AddUserToWorkspace)
	g.PUT("/:userId", h.UpdateUserRole)
	g.DELETE("/:userId", h.RemoveUserFromWorkspace)
}

func (h *UserWorkspaceRoleHandler) GetUsersByWorkspaceID(c echo.Context) error {
	workspaceId, err := strconv.Atoi(c.Param("workspaceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid workspace ID"})
	}

	users, err := h.repo.GetUsersByWorkspaceID(uint(workspaceId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserWorkspaceRoleHandler) AddUserToWorkspace(c echo.Context) error {
	workspaceId, err := strconv.Atoi(c.Param("workspaceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid workspace ID"})
	}

	var userWorkspaceRole model.UserWorkspaceRole
	if err := c.Bind(&userWorkspaceRole); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	userWorkspaceRole.WorkspaceID = uint(workspaceId)
	if err := h.repo.Create(&userWorkspaceRole); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userWorkspaceRole)
}

func (h *UserWorkspaceRoleHandler) UpdateUserRole(c echo.Context) error {
	workspaceId, err := strconv.Atoi(c.Param("workspaceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid workspace ID"})
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user ID"})
	}

	var userWorkspaceRole model.UserWorkspaceRole
	if err := c.Bind(&userWorkspaceRole); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}

	userWorkspaceRole.UserID = uint(userId)
	userWorkspaceRole.WorkspaceID = uint(workspaceId)

	if err := h.repo.UpdateRole(&userWorkspaceRole); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, userWorkspaceRole)
}

func (h *UserWorkspaceRoleHandler) RemoveUserFromWorkspace(c echo.Context) error {
	workspaceId, err := strconv.Atoi(c.Param("workspaceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid workspace ID"})
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user ID"})
	}

	if err := h.repo.Delete(uint(userId), uint(workspaceId)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User removed from workspace"})
}
