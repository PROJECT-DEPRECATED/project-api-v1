package routes

import (
	"errors"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(app *gin.Engine) {
	Index(app)
	APIV1(app)
	Resources(app)
	Mirror(app)
}

func printErr(ctx *gin.Context, status int, err error) {
	log.Errorln(err)
	ctx.JSON(status, gin.H{"status": 404, "err": err.Error()})
}

func NotFound(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	if err != mongo.ErrNoDocuments {
		return InternlServerErrHandler(ctx, err)
	}

	printErr(ctx, 404, err)
	return true
}

func NoContentHandler(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	printErr(ctx, 204, err)
	return true
}

func UnauthorizedHandler(ctx *gin.Context, auth bool) bool {
	if auth {
		return false
	}

	err := errors.New("unauthorized request detected")
	log.Errorln(err)
	ctx.JSON(401, gin.H{"status": 401, "auth": auth, "err": err.Error()})
	return true
}

func InternlServerErrHandler(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	printErr(ctx, 500, err)
	return true
}
