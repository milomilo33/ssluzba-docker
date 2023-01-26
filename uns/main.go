package main

import (
	"fmt"
	"uns/controllers"
	"uns/database"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	// Initialize Database

	database.Connect("postgresql://pg/uns-db?user=postgres&password=root")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8003")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/uns")
	{
		api.POST("/student", controllers.RegisterStudent)
		api.POST("/profesor", controllers.RegisterProfesor)
	}
	return router
}
