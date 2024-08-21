package main

import (
	"framework/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message" : "hello world",
		})
	})

	router.Start(config.APP_PORT)
}