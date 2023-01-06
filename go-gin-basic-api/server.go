package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sdqwec/go-gin-basic-api/controller"
	"github.com/sdqwec/go-gin-basic-api/service"
)

var (
	userService    service.UserService       = service.New()
	UserController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()

	server.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.FindAll())
	})

	server.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.Save(ctx))
	})
	server.Run(":8080")
}
