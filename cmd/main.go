package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "OK",
			})
		})
	}
	r.Run()
}
