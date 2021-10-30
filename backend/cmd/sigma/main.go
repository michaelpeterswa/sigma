package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/michaelpeterswa/sigma/backend/cmd/util"
)

func main() {
	fmt.Println("sigma")
	util.Info()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
