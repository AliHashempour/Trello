package handler

import (
	"Trello/internal/model"
	"Trello/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Workspace struct {
	repo repository.Workspace
}

func NewWorkspace(repo repository.Workspace) *Workspace {
	return &Workspace{repo: repo}
}

func (h *Workspace) Register(g *echo.Group) {
	g.GET("/", h.GetWorkspaceList)
	g.GET("/:id", h.GetWorkspace)
	g.POST("/", h.CreateWorkspace)
	g.PUT("/:id", h.UpdateWorkspace)
	g.DELETE("/:id", h.DeleteWorkspace)
}

func (h *Workspace) GetWorkspaceList(c echo.Context) error {
	workspaceList, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"workspaceList": workspaceList})
}

func (h *Workspace) GetWorkspace(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.ErrBadRequest
	}

	workspace, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"workspace": workspace})
}

func (h *Workspace) CreateWorkspace(c echo.Context) error {
	var workspace model.Workspace
	if err := c.Bind(&workspace); err != nil {
		return echo.ErrBadRequest
	}
	if workspace.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "name is required"})
	}
	if err := h.repo.Create(&workspace); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"created": workspace})
}

func (h *Workspace) UpdateWorkspace(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.ErrBadRequest
	}
	var workspace model.Workspace
	if err := c.Bind(&workspace); err != nil {
		return echo.ErrBadRequest
	}
	if workspace.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "name is required"})
	}
	workspace.ID = uint(id)
	if err := h.repo.Update(&workspace); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"updated": workspace})
}

func (h *Workspace) DeleteWorkspace(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "delete workspace " + c.Param("id")})
}
