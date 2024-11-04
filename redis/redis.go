package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type user = struct{
	name string `json:"name"`
}

//　汎用repository interface
type repository interface{
	New() 
	Create(redisKey string, user user)
	Get(redisKey string)(string, error)
}

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

func(r *redisRepo) Create(redisKey string, user user){
	jsonData, err := json.Marshal(user)
	if err !=nil{
		fmt.Println("fail to create redis data")
	}
	r.redisConn.Set(context.Background(), redisKey, jsonData, 0)
}