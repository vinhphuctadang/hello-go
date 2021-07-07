package main

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
	name := c.DefaultQuery("name", "NO_NAME")

	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

func getPerson(c *gin.Context) {
	personId := c.Param("id")

	c.JSON(200, gin.H{
		"id":   personId,
		"name": "John Levine",
		"age":  32,
	})
}

func main() {
	r := gin.Default() //gin.Default() creates a Gin router with default middleware: logger and recovery middleware.
	r.GET("/ping", ping)
	r.GET("/hello", hello)
	r.GET("/person/:id", getPerson)
	r.Run(":3000") // serve on port 3000
}
