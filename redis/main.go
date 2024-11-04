package main

import (
	"redis_test/middleware"

	"github.com/gin-gonic/gin"
)
var cookieKey = "cookieKey"

func main() {
	r := gin.Default()

	redisRepo := &redisRepo{}
	redisRepo.New()
	loginCheckGroup := r.Group("/", middleware.Checklogin(redisRepo))
	{
		loginCheckGroup.GET("/", HandlerGet)
		
	}
	r.Run(":3000")

}
