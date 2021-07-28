package main

import (
    "fmt"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(200, "Hello, Worldd!sheewshhh")
}

func main() {
	r := gin.Default()

	r.GET("/hello", HelloWorld)

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	err := r.Run(":3000")
	
	if err != nil {
		fmt.Print("linter, dont move!!")
	}
}
