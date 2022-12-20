package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type EditMode string

const (
	Name     EditMode = "name"
	Password EditMode = "password"
	IsOwner  EditMode = "is_owner"
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
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: a.Email}}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (a *Account) SetAccount(mode EditMode) (*Account, error) {
	var data Account
	coll := DB.Collection("account")
	var update bson.D
	filter := bson.D{{Key: "email", Value: a.Email}}

	switch mode {
	case Name:
		update = bson.D{{Key: string(Name), Value: a.Name}}
	case Password:
		update = bson.D{{Key: string(Password), Value: a.Password}}
	case IsOwner:
		update = bson.D{{Key: string(IsOwner), Value: a.IsOwner}}
	default:
		return nil, fmt.Errorf("invalid mode: %s", string(mode))
	}

	_, err := coll.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}

	return &data, nil
}
