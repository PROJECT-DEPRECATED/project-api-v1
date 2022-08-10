package api

import (
	"context"
	"strconv"

	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type LedData struct {
	Red   int `bson:"red"`
	Green int `bson:"green"`
	Blue  int `bson:"blue"`
}

func GetLed(c *gin.Context) {
	conf, _ := config.Get()

	coll := utils.DB.Database(conf.Database.DbName).Collection("led_data")
	res := coll.FindOne(context.TODO(), bson.D{{}})
	if res.Err() != nil {
		c.JSON(502, gin.H{"status": "502"})
		return
	}
}

func SetLed(c *gin.Context) {
	conf, _ := config.Get()
	red, err := strconv.ParseInt(c.PostForm("red"), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       "500",
			"reason":       err.Error(),
			"insert_value": c.PostForm("red"),
		})

		return
	}

	green, err := strconv.ParseInt(c.PostForm("green"), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       "500",
			"reason":       err.Error(),
			"insert_value": c.PostForm("green"),
		})

		return
	}

	blue, err := strconv.ParseInt(c.PostForm("blue"), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       "500",
			"reason":       err.Error(),
			"insert_value": c.PostForm("blue"),
		})

		return
	}

	if (int(red) > 255 || int(red) < 0) || (int(green) > 255 || int(green) < 0) || (int(blue) > 255 || int(blue) < 0) {
		c.JSON(500, gin.H{
			"status": "500",
			"reason": "please insert value 0 between 255.",
		})

		return
	}

	coll := utils.DB.Database(conf.Database.DbName).Collection("led_data")
	res := coll.FindOneAndUpdate(context.TODO(), bson.D{{Key: "_id", Value: 0}}, bson.D{
		{Key: "red", Value: red},
		{Key: "green", Value: green},
		{Key: "blue", Value: blue},
	})
	if res.Err() != nil {
		c.JSON(502, gin.H{"status": "502"})
		return
	}

	c.JSON(200, gin.H{
		"status": "200",
		"type":   "POST",
		"red":    int(red),
		"green":  int(green),
		"blue":   int(blue),
	})
}
