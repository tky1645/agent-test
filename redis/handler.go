package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandlerGet(c *gin.Context) {
	fmt.Println("call handler")
}