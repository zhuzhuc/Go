package main

import "github.com/gin-gonic/gin"

func main() {
	// Obtain the router instance
	r := gin.Default()

	// Define a route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "%v", "zzc")
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "%v", id)
	})
	// Start the HTTP server
	r.Run()
}
