package api

import (
	"context"
	"fmt"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-website/utils"
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

// Checking post is exist
func (p Post) isExist() bool {
	dummy := Post{}
	filter := bson.D{{Key: "_id", Value: p.ID}}
	if err := postColl().FindOne(context.TODO(), filter).Decode(&dummy); err != nil {
		return false
	}

	return true
}

// Show all posts list
func query(cur *mongo.Cursor) ([]*Post, error) {
	var results []*Post
	if err := cur.All(context.TODO(), &results); err != nil {
		fmt.Println("test")
		return nil, err
	}

	return results, nil
}

// Added new post
func (p Post) AddPost() error {
	res, err := postColl().InsertOne(context.TODO(), p)
	if err != nil {
		return err
	}

	log.Debugf("Document inserted with ID: %d\n", res.InsertedID)
	return nil
}

// Deleted exist post
func (p Post) DropPost() error {
	if !p.isExist() {
		return mongo.ErrNoDocuments
	}

	filter := bson.D{{Key: "_id", Value: p.ID}}
	res, err := postColl().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	log.Debugf("Document deleted: %d\n", res.DeletedCount)
	return nil
}

// Getting exist post
func (p Post) GetPost() (*Post, error) {
	var data Post
	filter := bson.D{{Key: "_id", Value: p.ID}}
	if err := postColl().FindOne(context.TODO(), filter).Decode(&data); err != nil {
		return nil, err
	}

	log.Debugf("Document finded with ID: %d\n", p.ID)
	return &data, nil
}

// Set exist post's content
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

// Find target post
func (p Post) SearchPost() ([]*Post, error) {
	filter := bson.D{{Key: "title", Value: p.Title}}
	res, err := postColl().Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	items, err := query(res)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Get all posts
func GetPosts() ([]*Post, error) {
	res, err := postColl().Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	items, err := query(res)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Get all posts length
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

// Setting post's id
func DefineID() (int, error) {
	var id int = -1
	cnt, err := GetPosts()
	if err != nil {
		return id, err
	}

	if cnt == nil {
		return 1, nil
	}

	for i, j := range cnt {
		def := i + 2
		if j == nil {
			id = def
			return id, nil
		}

		id = def
	}
	return id, nil
}
