package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", saludar)
	e.GET("/dividir", dividir)

	e.Start(":8080")
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)

	if f == 0 {
		return c.String(http.StatusBadRequest, "El dividendo no puede ser 0")
	}
	r := 3000 / f

	return c.String(http.StatusOK, strconv.Itoa(r))
}
