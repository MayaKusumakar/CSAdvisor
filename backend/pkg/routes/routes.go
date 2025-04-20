package routes

import (
	"SwipeNGo/pkg/mongoApi"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//[{
//event_id: string,
//owner_id: string,
//owner_name: string,
//startTime: string,
//endTime: string,
//location: {
//latitude: float,
//longitude: float,
//address: string
//},
//title: string,
//description: string,
//num_attending; int,
//banner_url: "https://unsplash.it/400/200"
//}]

type eventReturn struct {
	Event_id      string              `json:"event_id"`
	Owner_id      string              `json:"owner_id"`
	Owner_name    string              `json:"owner_name"`
	Start_time    string              `json:"start_time"`
	End_time      string              `json:"end_time"`
	Location      eventLocationReturn `json:"location"`
	Title         string              `json:"title"`
	Description   string              `json:"description"`
	Num_attending int                 `json:"num_attending"`
	Banner_url    string              `json:"banner_url"`
}

type eventLocationReturn struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

func Init(router *gin.Engine) {
	router.GET("/getEvents", func(c *gin.Context) {
		events, err := mongoApi.GetEvents()
		jsonArr := []eventReturn{}
		for _, event := range events {
			jsonArr = append(jsonArr, eventReturn{
				Event_id:   "test-event",
				Owner_id:   event.OwnerId,
				Owner_name: event.OwnerName,
				Start_time: event.StartTime,
				End_time:   event.EndTime,
				Location: eventLocationReturn{
					Latitude:  event.Location.Latitude,
					Longitude: event.Location.Longitude,
					Address:   event.Location.Address,
				},
				Title:         event.Title,
				Description:   event.Description,
				Num_attending: event.NumAttending,
				Banner_url:    "https://unsplash.it/400/200",
			})
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			jsonData, _ := json.Marshal(jsonArr)
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   string(jsonData),
			})
		}
	})
}
