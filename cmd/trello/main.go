package main

import (
	"Trello/internal/database"
	"Trello/internal/http/handler"
	"Trello/internal/http/middleware"
	"Trello/internal/repository"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.InfoLogger)
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	workspaceGroup := e.Group("/workspace")
	workspaceHandler := handler.NewWorkspace(repository.NewWorkspaceRepo(db))
	workspaceHandler.Register(workspaceGroup)

	userGroup := e.Group("/user")
	userHandler := handler.NewUser(repository.NewUserRepo(db))
	userHandler.Register(userGroup)

	userWorkspaceRoleGroup := e.Group("/workspace/:workspaceId/users")
	userWorkspaceRoleHandler := handler.NewUserWorkspaceRoleHandler(repository.NewUserWorkspaceRepository(db))
	userWorkspaceRoleHandler.Register(userWorkspaceRoleGroup)

	e.Logger.Fatal(e.Start(":8080"))

}
