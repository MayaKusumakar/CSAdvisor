package main

import (
	"github.com/gin-gonic/gin"
	"io"
)

func main() {

	// 	mongoClient, _ := mongoApi.GetMongoClient()
	// 	collection, err := mongoApi.GetCollection(mongoClient, "Some", "Thing")
	// 	fmt.Println(collection, err)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/ping", func(c *gin.Context) {
		jsonData, _ := io.ReadAll(c.Request.Body)
		c.JSON(200, gin.H{
			"message":      "pong",
			"message_body": string(jsonData),
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
