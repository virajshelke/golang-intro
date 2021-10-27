package main

import "github.com/gin-gonic/gin"

// 'go get -u github.com/gin-gonic/gin'
// or else add it here and use 'go mod tidy'

func main() {
	// creating instance of engine with the Logger and Recovery middleware already attached.
	r := gin.Default()

	// creating a route and passing a callback function
	r.GET("/ping", func(c *gin.Context) {
		// using the context instance passed inside the callback to send the response
		// gin.H is a shortcut for creating a map[string]interface{}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// attaches the router to a http.Server and starts listening and serving HTTP requests
	r.Run()
}
