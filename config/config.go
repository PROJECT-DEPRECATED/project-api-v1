package config

import (
	"encoding/json"
	"os"
)

var (
	filename  = "./config.json"
	GetSample = `{
		"token": "SPECTIFY_TOKEN",
		"password_salt": "PASSWORD_SALT",
		"hangang_api": {
			"url": "HANGANG_DATA_URL"
		},
		"mojang_api": {
			"api_url": "MOJANG_API_URL",
			"session_url": "SESSION_URL"
		},
		"database": {
			"url": "MONGO_DATABASE_URL",
			"port": 27017,
			"db_name": "projecttl-website",
			"username": "MONGO_DATABASE_USERNAME",
			"password": "MONGO_DATABASE_PASSWORD"
		}
	}`
)

type conf struct {
	Token        string `json:"token"`
	PasswordSalt string `json:"password_salt"`
	HangangAPI   struct {
		Url string `json:"url"`
	} `json:"hangang_api"`
	MojangAPI struct {
		API     string `json:"api_url"`
		Session string `json:"session_url"`
	} `json:"mojang_api"`
	Database struct {
		Url      string `json:"url"`
		Port     int    `json:"port"`
		DbName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
}

func Get() (*conf, error) {
	config, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data conf
	err = json.Unmarshal(config, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
