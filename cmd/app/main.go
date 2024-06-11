package main

import (
	"Trello/internal/http/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.InfoLogger)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.Logger.Fatal(e.Start(":8080"))

}
