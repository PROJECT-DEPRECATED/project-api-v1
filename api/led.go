package api

import (
	context2 "context"

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

func GetLed(context *gin.Context) {
	conf, _ := config.Get()
	coll := utils.DB.Database(conf.Database.DbName).Collection("led_data")
	res := coll.FindOne(context2.TODO(), bson.D{{}})
	if res.Err() != nil {
		context.JSON(502, gin.H{"status": "502"})
		return
	}
}

func SetLed(context *gin.Context) {

}
