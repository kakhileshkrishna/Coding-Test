package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Temporary route for testing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Go backend!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
