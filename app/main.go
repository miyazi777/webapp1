package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := Router()

	e.Logger.Fatal(e.Start(":3030"))
}

func Router() *echo.Echo {
	e := echo.New()
	initRouting(e)
	return e
}

func initRouting(e *echo.Echo) {
	e.GET("/", count)
}

func count(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]int{"count": 42})
}
