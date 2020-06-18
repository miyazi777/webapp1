package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	initRouting(e)

	e.Logger.Fatal(e.Start(":1313"))
}

func initRouting(e *echo.Echo) {
	e.GET("/", count)
}

func count(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]int{"count": 42})
}
