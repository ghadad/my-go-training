package main

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"slice":   make([]string, 10),
	})
}

func Handler2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"slice":   make([]string, 10),
	})
}

func Handler3(c *gin.Context) {

	c.JSON(200, gin.H{
		"version": "v1",
		"message": "pong",
		"slice":   make([]string, 10),
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", Handler)
	r.GET("/", Handler2)
	r.GET("/v1", Handler3)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
