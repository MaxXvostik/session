package main

import (
	"example.com/sessionRedisGo/src/controller"
	"example.com/sessionRedisGo/src/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp",
		"localhost:6379", "", []byte("secret"))

	router.Use(sessions.Sessions("mysession", store))

	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)

	auth := router.Group("/auth")
	auth.Use(middleware.Authentication())

	auth.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Everything is ok",
		})
	})

	router.Run(":8000")

}
