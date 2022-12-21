package api

import (
	"context"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-website/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	ID      int    `bson:"_id"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Created string `bson:"created"`
}

func postColl() *mongo.Collection {
	return utils.DB.Collection("post")
}

func (p Post) AddPost() error {
	res, err := postColl().InsertOne(context.TODO(), p)
	if err != nil {
		return err
	}

	log.Debugf("Document inserted with ID: %d\n", res.InsertedID)
	return nil
}

func (p Post) DropPost() error {
	filter := bson.D{{Key: "_id", Value: p.ID}}
	res, err := postColl().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	log.Debugf("Document deleted: %d\n", res.DeletedCount)
	return nil
}

func (p Post) GetPost() (*Post, error) {
	var data *Post
	filter := bson.D{{Key: "_id", Value: p.ID}}
	if err := postColl().FindOne(context.TODO(), filter).Decode(data); err != nil {
		return nil, err
	}

	log.Debugf("Document finded with ID: %d\n", p.ID)
	return data, nil
}

func (p Post) SetPost() error {
	filter := bson.D{{Key: "_id", Value: p.ID}}
	update := bson.D{{Key: "$set", Value: p}}
	res, err := postColl().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	log.Debugf("Documents updated: %v\n", res.ModifiedCount)
	return nil
}

func (p Post) SearchPost() ([]gin.H, error) {
	var results []Post
	filter := bson.D{{Key: "title", Value: p.Title}}
	res, err := postColl().Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}

	if err = res.All(context.TODO(), results); err != nil {
		return nil, err
	}

	var items []gin.H
	for _, i := range results {
		res.Decode(&i)
		items = append(items, gin.H{
			"_id":     i.ID,
			"title":   i.Title,
			"content": i.Content,
			"created": i.Created,
		})
	}

	return items, nil
}

func GetPosts() ([]*Post, error) {
	var all []*Post
	var data *Post
	res, err := postColl().Find(context.TODO(), Post{})
	if err != nil {
		return nil, err
	}

	for res.Next(context.TODO()) {
		err = res.Decode(&data)
		if err != nil {
			return nil, err
		}

		all = append(all, data)
	}

	return all, nil
}

func GetPostsCount() (int, error) {
	count := 0
	posts, err := GetPosts()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return count, nil
		}

		return -1, err
	}

	for i, j := range posts {
		if j == nil {
			break
		}

		count = i
	}

	return count, nil
}

func DefineID() (int, error) {
	var id int = -1
	cnt, err := GetPosts()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			id = 1
			return id, nil
		}

		return id, err
	}

	for i, j := range cnt {
		if j == nil {
			id = i
			return id, nil
		}

		id = i
	}

	return id, nil
}
