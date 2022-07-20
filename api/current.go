package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CurrentDate struct {
	TimeZone string `json:"timezone"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Date     int    `json:"date"`
	Time     struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
		Second int `json:"second"`
	} `json:"time"`
}

func getCurrent() CurrentDate {
	time := time.Now()
	zone, _ := time.Zone()

	current := CurrentDate{
		TimeZone: zone,
		Year:     time.Year(),
		Month:    int(time.Month()),
		Date:     time.Day(),
		Time: struct {
			Hour   int "json:\"hour\""
			Minute int "json:\"minute\""
			Second int "json:\"second\""
		}{
			Hour:   time.Hour(),
			Minute: time.Minute(),
			Second: time.Second(),
		},
	}

	return current
}

func CurrentTime(context *gin.Context) {
	context.JSON(200, gin.H{
		"data":   getCurrent(),
		"status": "200",
	})
}
