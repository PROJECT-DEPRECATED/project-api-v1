package routes

import (
	"errors"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NotFound(ctx *gin.Context, err error) bool {
	if err != nil {
		return false
	}

	if err != mongo.ErrNoDocuments {
		return InternlServerErrHandler(ctx, err)
	}

	log.Errorln(err)
	ctx.JSON(204, gin.H{"status": 404, "err": err.Error()})
	return true
}

func NoContentHandler(ctx *gin.Context, err error) bool {
	if err != nil {
		return false
	}

	log.Errorln(err)
	ctx.JSON(204, gin.H{"status": 204, "err": err.Error()})
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

	log.Errorln(err)
	ctx.JSON(500, gin.H{"status": 500, "err": err.Error()})
	return true
}
