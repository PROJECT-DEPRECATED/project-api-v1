package utils

import (
	"context"
	"errors"

	"github.com/devproje/project-website/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	UniqueId string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (a *Account) isExist() bool {
	var data Account

	conf, _ := config.Get()
	coll := DB.Database(conf.Database.DbName).Collection("account")
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: a.Email}}).Decode(&data)

	return err == nil
}

func (a *Account) New() error {
	if a.isExist() {
		return errors.New("email already exist")
	}
	conf, _ := config.Get()
	coll := DB.Database(conf.Database.DbName).Collection("account")
	_, err := coll.InsertOne(context.TODO(), bson.D{
		{Key: "_id", Value: uuid.NewString()},
		{Key: "name", Value: a.Name},
		{Key: "email", Value: a.Email},
		{Key: "password", Value: a.Password},
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Drop() error {
	if !a.isExist() {
		return errors.New("email not exist")
	}
	conf, _ := config.Get()
	coll := DB.Database(conf.Database.DbName).Collection("account")
	_, err := coll.DeleteOne(context.TODO(), bson.D{{Key: "email", Value: a.Email}})
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Find() (*Account, error) {
	var data Account

	conf, _ := config.Get()
	coll := DB.Database(conf.Database.DbName).Collection("account")
	err := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: a.Name}}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
