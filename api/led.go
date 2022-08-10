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
	Red   int32 `bson:"red"`
	Green int32 `bson:"green"`
	Blue  int32 `bson:"blue"`
}

func GetLed(c *gin.Context) {
	conf, _ := config.Get()
	err := utils.AuthUtils(c)
	if err != nil {
		return
	}

	coll := utils.DB.Database(conf.Database.DbName).Collection("led_data")
	var result LedData
	res := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: 0}})
	if res.Err() != nil {
		c.JSON(500, gin.H{
			"status": "500",
			"reason": res.Err().Error(),
		})
		return
	}
	defer res.Decode(&result)

	c.JSON(200, gin.H{
		"status": "200",
		"type":   "GET",
		"red":    int(result.Red),
		"green":  int(result.Green),
		"blue":   int(result.Blue),
	})
}

func SetLed(c *gin.Context) {
	conf, _ := config.Get()
	err := utils.AuthUtils(c)
	if err != nil {
		return
	}

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
	filter := bson.D{{Key: "_id", Value: 0}}
	docs := bson.M{"$set": bson.D{
		{Key: "red", Value: red},
		{Key: "green", Value: green},
		{Key: "blue", Value: blue},
	}}
	res := coll.FindOneAndUpdate(context.TODO(), filter, docs)
	if res.Err() != nil {
		c.JSON(500, gin.H{
			"status": "500",
			"reson":  res.Err().Error(),
		})
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
