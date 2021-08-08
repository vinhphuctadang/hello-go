package main

import (
	"fmt"
	"http-service-example/internal/database"
	"http-service-example/internal/handlers"
	"github.com/gin-gonic/gin"
)

func setupServerRoutes(server *gin.Engine) {
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	dataAdapter, err := database.NewDataAdapter()
	if err != nil {
		panic(err)
	}

	identity, err := dataAdapter.GetUserByUsernameAndPassword("phuc", "haha")
	fmt.Println(identity)
	authenticate := &handlers.Authenticate {
		Adapter: dataAdapter,
	}

	server.POST("/auth", authenticate.Handler())
}

func main() {
	server := gin.Default()
	setupServerRoutes(server)
	server.Run(":8080")
}