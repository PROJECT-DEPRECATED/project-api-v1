package utils

import (
	"context"
	"fmt"

	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() {
	var err error
	conf, _ := config.Get()
	data := conf.Database

	uri := fmt.Sprintf("mongodb://%s:%d/%s", data.Url, data.Port, data.DbName)
	option := options.Client().ApplyURI(uri)
	option.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      data.Username,
		Password:      data.Password,
	})
	DB, err = mongo.Connect(context.TODO(), option)
	if err != nil {
		log.Logger.Errorln("Database connection failed.")
		log.Logger.Errorln(err)
		return
	}

	log.Logger.Infoln("Database connected.")
}
