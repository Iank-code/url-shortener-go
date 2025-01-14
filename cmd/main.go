package main

import (
	"github.com/deatil/go-encoding/encoding"
	"github.com/gin-gonic/gin"
)

func encodeBase62(data string) string {
	base64Data := encoding.FromString(data).Base62Encode().ToString()

	return base64Data
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200,gin.H{
			"message": "Pong",
			"base62Encoded": encodeBase62("https:smartplyconnect.com"),
		})
	})

	r.Run()
}
