package main

import (
	"DDD/command/user"
	"DDD/query/plant"

	"github.com/gin-gonic/gin"
)


func main (){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ping pong")
	})
	r.GET("/users", user.HandlerGET)
	r.POST("/users", user.HandlerPOST)
	r.PUT("/users", user.HandlerPUT)

	r.POST("/plants", plant.HandlerPOST)
	r.PATCH("/plants/:id", plant.HandlerPATCH)
	r.Run(":8080")

}