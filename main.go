package main

import (
	"gin-graphql/graph"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(cors.Default())
	router.POST("/todo", graph.TodoGraphRouter)

	// Ping test
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
