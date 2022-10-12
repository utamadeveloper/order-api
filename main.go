package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/utamadeveloper/golang-rest-api/config"
	"github.com/utamadeveloper/golang-rest-api/handlers"
	"github.com/utamadeveloper/golang-rest-api/middleware"
)

var APP_HOST = "0.0.0.0"
var APP_PORT = 8000

func main() {
	// Database connect
	config.PostgresConnect()

	// Router
	r := gin.Default()
	r.Use(middleware.JSONResponseMiddleware())

	// User
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.FindAllUser)
	r.GET("/users/:id", handlers.FindOneUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// Order
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.FindAllOrder)
	r.GET("/orders/:id", handlers.FindOneOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)
	r.GET("/orders/user/:user_id", handlers.FindAllOrderUser)

	APP_SERVER := fmt.Sprintf("%s:%d", APP_HOST, APP_PORT)
	r.Run(APP_SERVER)
}
