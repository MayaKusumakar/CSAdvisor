package main

import (
	"SwipeNGo/pkg/routes"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {

	//arr, _ := mongoApi.GetEvents()
	//fmt.Println((arr[1]))
	//newEvent := mongoApi.Event{
	//	"",
	//	"Owner Three",
	//	"Jane Doe The Second",
	//	"An Event",
	//	mongoApi.Location{
	//		112.23,
	//		222.345,
	//		"12 S street, Davis, CA",
	//	},
	//	"time started",
	//	"time ended",
	//	"Some description here",
	//	73,
	//}
	//
	//client, _ := mongoApi.GetMongoClient()
	//collection, _ := mongoApi.GetCollection(client, "main", "events")
	//err := mongoApi.AddEvent(collection, newEvent)
	//if err != nil {
	//	log.Fatal(err)
	//}

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
	routes.Init(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
