package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!echo!")
	})
	e.Static("/", "static/")
	e.Logger.Fatal(e.Start(":8080"))
}
