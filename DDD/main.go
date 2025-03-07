package main

import (
	"DDD/command/user"
	"DDD/query/plant"
	"database/sql"
	"fmt"

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
		DBName: "sampledb",    // Use environment variables
	}
	dsn := cfg.FormatDSN() // Automatically generate the correct DSN

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful")

	// Initialize user handlers
	user.InitHandlers(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ping pong")
	})
	r.GET("/users", user.HandlerGET)
	r.POST("/users", user.HandlerPOST)
	r.PUT("/users/:id", user.HandlerPUT)
	r.GET("/users/:id", user.HandlerFETCH)
	r.DELETE("/users/:id", user.HandlerDELETE)

	r.POST("/plants", plant.HandlerPOST)
	r.PATCH("/plants/:id", plant.HandlerPATCH)

	r.Run(":8080")
}
