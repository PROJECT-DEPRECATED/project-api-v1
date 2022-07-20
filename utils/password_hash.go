package utils

import (
	"github.com/devproje/project-website/config"
	"github.com/stretchr/objx"
)

func PasswordHash(password string) *string {
	hash := objx.HashWithKey(password, config.Get().PasswordSalt)
	return &hash
}
