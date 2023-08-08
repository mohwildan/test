package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "apa kata mu mus ti lorem ipsum dolor si amet mantap cuy",
			"data": map[string]any{
				"name":  "wildan",
				"email": "mohwildanwildan15@gmail.com",
			},
		})
	})
	r.Run(":3002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
