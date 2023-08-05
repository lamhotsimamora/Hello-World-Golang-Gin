package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World! My first website developed by Golang",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "home",
		})
	})
	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "about",
		})
	})
	r.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "profile",
		})
	})
	r.GET("/contact", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "contact",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
