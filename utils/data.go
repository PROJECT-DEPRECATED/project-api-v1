package utils

import (
	"context"
	"fmt"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-website/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() error {
	var err error
	var client *mongo.Client
	conf, _ := config.Get()
	data := conf.Database

	uri := fmt.Sprintf("mongodb://%s:%d", data.URL, data.Port)
	option := options.Client().ApplyURI(uri)
	option.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      data.Username,
		Password:      data.Password,
	})
	client, err = mongo.Connect(context.TODO(), option)
	if err != nil {
		log.Errorln("Database connection failed.")
		return err
	}

	DB = client.Database(data.DbName)
	log.Infoln("Database connected.")
	return nil
}
