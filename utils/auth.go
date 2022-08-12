package utils

import (
	"errors"

	"github.com/devproje/project-website/config"
	"github.com/gin-gonic/gin"
)

func isAuth(token string) bool {
	conf, _ := config.Get()
	return PasswordHash(token) == PasswordHash(conf.Token)
}

func AuthUtils(c *gin.Context) error {
	if !isAuth(c.GetHeader("Authorization")) {
		c.JSON(401, gin.H{"status": "401"})
		return errors.New("unauthorized")
	}

	return nil
}
