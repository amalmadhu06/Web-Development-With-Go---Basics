package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	//creating a router
	r := gin.Default()

	//adding a route and handler function
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!!!")
	})

	//starting the server
	r.Run("localhost:8080")
}
