package main

import (
	"DDD/command/user"
	"DDD/query/plant"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)


func main (){
	r := gin.Default()
// ハンドラーのインスタンス化
// データベース接続
cfg := mysql.Config{
    User:   "sampleuser",
    Passwd: "samplepass",
    Net:    "tcp",
    Addr:   "localhost:3306",
    DBName: "Watering",
}
dsn := cfg.FormatDSN() // 自動的に正しいDSNを生成

db, err := sql.Open("mysql", dsn)
if err != nil{
	panic(err)
}

if err := db.Ping(); err!=nil{
	panic(err)
}

fmt.Printf("no error")


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ping pong")
	})
	r.GET("/users", user.HandlerGET)
	r.POST("/users", user.HandlerPOST)
	r.PUT("/users", user.HandlerPUT)

	r.POST("/plants", plant.HandlerPOST)
	r.PATCH("/plants/:id", plant.HandlerPATCH)
	r.Run(":8080")

	//net/httpを使う場合
	// http.HandleFunc("/ping", func(http.ResponseWriter, *http.Request){
	// 	fmt.Println("ping pong")
	// })
	// http.ListenAndServe(":18080", nil)
}