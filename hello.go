package main

import (
    "fmt"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(200, "Hello, Worlddd2!make change")
}

func main() {
	r := gin.Default()

	r.GET("/", HelloWorld)

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// r.Use(static.Serve("/hello", static.LocalFile("./views", true)))

	err := r.Run(":3000")
	
	if err != nil {
		fmt.Print("linter, dont move!!")
	}
}
