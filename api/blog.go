package api

import (
	"context"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-website/utils"
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

func (p *Post) AddPost() error {
	id, err := GetPostsCount()
	if err != nil {
		return err
	}

	p.ID = id + 1

	res, err := postColl().InsertOne(context.TODO(), p)
	if err != nil {
		return err
	}

	log.Debugf("Inserted ID: %\n", res)
	return nil
}

func DropPost() {}

func GetPost() {}

func SetPost() {}

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
	cnt, err := GetPostsCount()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}

		return -1, err
	}
}
