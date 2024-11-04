package util

import (
	"github.com/gin-gonic/gin"
)

//　汎用repository interface
type Repository interface{
	New() 
	Create(c *gin.Context, redisKey string, user User)
	Get(redisKey string)(string, error)
	GetCount(redisKey string)(int, error)
}

type User = struct{
	Name string `json:"Name"`
	Nickname string `json:"nickName"`
	Count int `json:"count"`
}