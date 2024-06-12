package handler

import (
	"Trello/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
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
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "get workspace list"})
}
func (h *Workspace) GetWorkspace(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "get workspace " + c.Param("id")})
}
func (h *Workspace) CreateWorkspace(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "create workspace "})
}
func (h *Workspace) UpdateWorkspace(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "update workspace " + c.Param("id")})
}
func (h *Workspace) DeleteWorkspace(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "delete workspace " + c.Param("id")})
}
