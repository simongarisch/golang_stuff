package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default() // router
	r.LoadHTMLGlob("templates/*")

	// http://localhost:8080/
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "root",
		})
	})

	// http://localhost:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// http://localhost:8080/ginny
	r.GET("/ginny", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{},
		)

	})

	// http://localhost:8080/counting
	count := 0
	r.GET("/counting", func(c *gin.Context) {
		count++
		if count > 10 {
			count = 0
		}

		c.HTML(
			http.StatusOK,
			"counting.html",
			gin.H{"count": count},
		)

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
