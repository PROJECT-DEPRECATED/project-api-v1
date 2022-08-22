package utils

import (
	"context"
	"errors"

	"github.com/devproje/project-website/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	UniqueId string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (a *Account) isExist() bool {
	conf, _ := config.Get()
	data := Account{}
	coll := DB.Database(conf.Database.DbName).Collection("account")
	err := coll.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: a.Email},
	}).Decode(&data)

	return err == nil
}

func (a *Account) Add() error {
	if a.isExist() {
		return errors.New("this email is already exist")
	}
	conf, _ := config.Get()
	coll := DB.Database(conf.Database.DbName).Collection("account")
	_, err := coll.InsertOne(context.TODO(), bson.D{
		{Key: "_id", Value: a.UniqueId},
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
		return errors.New("this email isn't exist")
	}
	conf, _ := config.Get()
	coll := DB.Database(conf.Database.DbName).Collection("account")
	_, err := coll.DeleteOne(context.TODO(), bson.D{
		{Key: "email", Value: a.Email},
	})
	if err != nil {
		return err
	}

	return nil
}
