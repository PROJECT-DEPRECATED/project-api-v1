package utils

import (
	"errors"

	"github.com/devproje/project-website/config"
	"github.com/gin-gonic/gin"
)

func isAuth(token string) bool {
	conf, _ := config.Get()
	return SaltHash(token) == SaltHash(conf.Token)
}

func AuthUtils(c *gin.Context) error {
	if !isAuth(c.GetHeader("Authorization")) {
		return errors.New("unauthorized request detected")
	}

	return nil
}
