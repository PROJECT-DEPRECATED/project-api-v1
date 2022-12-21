package routes

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/devproje/project-website/api"
	"github.com/gin-gonic/gin"
)

func APIV1(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("/current", api.CurrentTime)
		v1.GET("/hangang", api.Hangang)
		v1.GET("/hangang/:area", api.Hangang)
		v1.GET("/mcprofile/:username", api.MCProfile)
		v1.GET("/led", api.GetLed)
		v1.POST("/led", api.SetLed)

		blog := v1.Group("/blog")
		{
			blogRouter(blog)
		}
	}
}

func blogRouter(group *gin.RouterGroup) {
	var getID = func(ctx *gin.Context) (int, error) {
		id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
		return int(id), err
	}

	group.POST("/add", func(ctx *gin.Context) {
		id, err := api.DefineID()
		if InternlServerErrHandler(ctx, err) {
			return
		}

		title := ctx.PostForm("title")
		content := ctx.PostForm("content")
		if title == "" || content == "" {
			NoContentHandler(ctx, errors.New("no content"))
			return
		}

		form := api.Post{
			ID:      id,
			Title:   title,
			Content: content,
			Created: time.Now().Format(time.RFC3339),
		}
		if form.AddPost(); InternlServerErrHandler(ctx, err) {
			return
		}

		ctx.JSON(200, gin.H{"status": 200, "type": "add", "_id": id})
	})
	group.POST("/drop/:id", func(ctx *gin.Context) {
		id, err := getID(ctx)
		if InternlServerErrHandler(ctx, err) {
			return
		}

		form := api.Post{ID: id}
		if err = form.DropPost(); InternlServerErrHandler(ctx, err) {
			return
		}

		ctx.JSON(200, gin.H{"status": 200, "type": "drop", "_id": id})
	})
	group.POST("/set/:id", func(ctx *gin.Context) {
		id, err := getID(ctx)
		if InternlServerErrHandler(ctx, err) {
			return
		}

		title := ctx.PostForm("title")
		content := ctx.PostForm("content")
		if title == "" || content == "" {
			NoContentHandler(ctx, errors.New("no content"))
			return
		}

		form := api.Post{
			ID:      id,
			Title:   title,
			Content: content,
		}
		if form.SetPost(); InternlServerErrHandler(ctx, err) {
			return
		}

		ctx.JSON(200, gin.H{"status": 200, "type": "set", "_id": id})
	})
	group.GET("/post/:id", func(ctx *gin.Context) {
		id, err := getID(ctx)
		if InternlServerErrHandler(ctx, err) {
			return
		}

		post, err := api.Post{ID: id}.GetPost()
		if NotFound(ctx, err) {
			return
		}

		ctx.JSON(200, gin.H{
			"status":  200,
			"type":    "get",
			"_id":     post.ID,
			"title":   post.Title,
			"content": post.Content,
			"created": post.Created,
		})
	})
	group.GET("/posts", func(ctx *gin.Context) {
		posts, err := api.GetPosts()
		if NotFound(ctx, err) {
			return
		}

		fmt.Println(posts)

		ctx.JSON(200, gin.H{"status": 200, "type": "query", "page": posts})
	})
	group.GET("/search/:title", func(ctx *gin.Context) {
		title := strings.ReplaceAll(ctx.Param("title"), "%20", " ")
		res, err := api.Post{Title: title}.SearchPost()
		if NotFound(ctx, err) {
			return
		}

		ctx.JSON(200, gin.H{"status": 200, "page": res})
	})
}
