package main

import (
	"DDD/command/user"
	"DDD/query/plant"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	// Database connection
	cfg := mysql.Config{
		User:   "sampleuser", // Use environment variables
		Passwd: "samplepass", // Use environment variables
		Net:    "tcp",
		Addr:   "ddd_rdb:3306", // Use environment variables
		DBName: "mydb",    // Use environment variables
	}
	dsn := cfg.FormatDSN() // Automatically generate the correct DSN

	var db *sql.DB
	var err error
	const maxRetries = 10
	for i := 0; i < maxRetries; i++ {
	db, err = sql.Open("mysql", dsn)
	if err == nil {
		err = db.Ping()
	}
	if err == nil {
		fmt.Println("Database connection successful")
		break
	}
	fmt.Printf("MySQL接続失敗 (%d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic(fmt.Sprintf("MySQLに接続できませんでした: %v", err))
	}
	defer db.Close()


	// Initialize user handlers
	user.InitHandlers(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ping pong")
	})
	r.GET("/users", user.HandlerGET)
	r.POST("/users", user.HandlerPOST)
	r.PUT("/users/:id", user.HandlerPUT)
	// r.GET("/users/:id", user.HandlerFETCH)
	// r.DELETE("/users/:id", user.HandlerDELETE)

	r.POST("/plants", plant.HandlerPOST)
	r.GET("/plants", plant.HandlerGETPlants)
	r.PATCH("/plants/:id", plant.HandlerPATCH)

	r.Run(":8080")
}
