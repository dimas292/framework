package main

import (
	"framework/config"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
	router := echo.New()

	router.Use(Logger)

	router.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message" : "hello world",
		})
	})

	userRouter := router.Group("user")
	{
		userRouter.GET("", func(c echo.Context) error {
			log.Println("log success from user")

			return c.JSON(http.StatusOK, echo.Map{
				"messsages" : "get from user",
			})
		})
	}

	router.Start(config.APP_PORT)
}

func Logger(next echo.HandlerFunc)echo.HandlerFunc{
	return func(c echo.Context) error {
		log.Println(c.Request().Method, c.Request().URL.Path)
		return next(c)
	}
}