package utils

import (
	"github.com/devproje/project-website/config"
	"github.com/stretchr/objx"
)

func PasswordHash(password string) *string {
	conf, _ := config.Get()
	hash := objx.HashWithKey(password, conf.PasswordSalt)
	return &hash
}
