package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"redis_test/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// redisのrepo
type redisRepo  struct{
	redisConn   *redis.Client
} 


func(r *redisRepo) New() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})
	r.redisConn = redisClient
}

func(r *redisRepo) Get(redisKey string)(string, error){
	redisValue, err:= r.redisConn.Get(context.Background(), redisKey).Result()
	if err != redis.Nil &&err !=nil{
		fmt.Println("fait to get redis data")
	}
	return redisValue, err
}
func(r *redisRepo) GetCount(redisKey string)(int, error){
	redisValue, err:= r.redisConn.Get(context.Background(), redisKey).Result()
	if err != redis.Nil &&err !=nil{
		fmt.Println("fait to get redis data")
	}
	var data util.User
	err = json.Unmarshal([]byte(redisValue), &data)
	if err != nil {
		fmt.Println("fait to unmershal redis data")
    }
	return data.Count, err
}

func(r *redisRepo) Create(c *gin.Context, redisKey string, user util.User){
	fmt.Printf("user before json mershal : %+v\n", user)
	fmt.Println()
	jsonData, err := json.Marshal(user)
	if err !=nil{
		fmt.Println("fail to create json data")
		return
	}
	fmt.Printf("jsondata : %s", string(jsonData))
	fmt.Println()
	status := r.redisConn.Set(c, redisKey, jsonData, time.Duration(10) * time.Second)
	if status.Err() != nil {
		fmt.Println("fail to set redis")
	}
}
