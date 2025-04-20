package main

import (
	"SwipeNGo/pkg/mongoApi"
	"fmt"
)

func main() {

	mongoClient, _ := mongoApi.GetMongoClient()
	collection, err := mongoApi.GetCollection(mongoClient, "Some", "Thing")
	fmt.Println(collection, err)

	//router := gin.Default()
	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//router.Run() // listen and serve on 0.0.0.0:8080
}
