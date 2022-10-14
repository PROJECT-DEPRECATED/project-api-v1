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

type Config struct {
	Token        string `json:"token"`
	PasswordSalt string `json:"password_salt"`
	HangangAPI   struct {
		URL string `json:"url"`
	} `json:"hangang_api"`
	MojangAPI struct {
		URL        string `json:"api_url"`
		SessionURL string `json:"session_url"`
	} `json:"mojang_api"`
	Database struct {
		URL  string `json:"url"`
		Port int    `json:"port"`
		DbName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
}

func Get() (*Config, error) {
	config, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data Config
	err = json.Unmarshal(config, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
