package utils

import (
	"context"
	"fmt"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-website/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() error {
	var err error
	conf, _ := config.Get()
	data := conf.Database

	uri := fmt.Sprintf("mongodb://%s:%d/%s", data.URL, data.Port, data.DbName)
	option := options.Client().ApplyURI(uri)
	option.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      data.Username,
		Password:      data.Password,
	})
	DB, err = mongo.Connect(context.TODO(), option)
	if err != nil {
		log.Errorln("Database connection failed.")
		return err
	}

	log.Infoln("Database connected.")
	return nil
}
