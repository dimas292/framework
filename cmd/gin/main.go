package main

import (
	"framework/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"log"
	"context"
)

func main() {

	router := gin.New()

	router.Use(Authorization())

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

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// do process validation
		// get user id from token
		userId := 10

		log.Println("Authrorization: set user id with", userId)
        // get context
		myCtx := ctx.Request.Context()
        // set context
		myCtx = context.WithValue(myCtx, "USER_ID", userId)
        // get request with new context
		req := ctx.Request.WithContext(myCtx)
        // change the request to new request with new context
		ctx.Request = req
		ctx.Next()

	}
}

