package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	UniqueId string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	IsOwner  bool   `bson:"is_owner"`
}

func (a *Account) isExist() bool {
	var data Account
	coll := DB.Collection("account")
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: a.Email}}).Decode(&data)

	return err == nil
}

func (a *Account) AddAccount() error {
	if a.isExist() {
		return fmt.Errorf("%s email already exist", a.Email)
	}
	coll := DB.Collection("account")
	_, err := coll.InsertOne(context.TODO(), bson.D{
		{Key: "_id", Value: uuid.New().String()},
		{Key: "name", Value: a.Name},
		{Key: "email", Value: a.Email},
		{Key: "password", Value: a.Password},
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) DropAccount() error {
	if !a.isExist() {
		return errors.New("email not exist")
	}
	coll := DB.Collection("account")
	_, err := coll.DeleteOne(context.TODO(), bson.D{{Key: "email", Value: a.Email}})
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) GetAccount() (*Account, error) {
	var data Account
	coll := DB.Collection("account")
	err := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: a.Name}}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (a *Account) SetAccount(mode string) (*Account, error) {
	var data Account
	coll := DB.Collection("account")
	var update bson.D
	filter := bson.D{{Key: "email", Value: a.Email}}

	switch mode {
	case "name":
		update = bson.D{{Key: "name", Value: a.Name}}
	case "password":
		update = bson.D{{Key: "password", Value: a.Password}}
	case "is_owner":
		update = bson.D{{Key: "is_owner", Value: a.IsOwner}}
	case "default":
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}

	_, err := coll.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}

	return &data, nil
}
