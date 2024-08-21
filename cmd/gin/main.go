package main

import (
	"framework/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"messages": "hellow",
			"version":  "no set",
		})
	})

	router.POST("/register", func(ctx *gin.Context) {

		var req interface{}

		auth := ctx.GetHeader("Authorization")
		name := ctx.Param("name")
		role := ctx.Query("role")

		err := ctx.ShouldBind(&req)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message" : "hello" + name, 
			"payload" : req,
			"header" : gin.H{
				"auth" : auth,
 			},
			"role" : gin.H{
				"role" : role,
			},
		})
	})

	router.GET("/users/:name", func(ctx *gin.Context) {
		
		name := ctx.Param("name")

		ctx.JSON(http.StatusOK, gin.H{
			"messages" : "hello" + name,
		})

	})

	router.Run(config.APP_PORT)

}
