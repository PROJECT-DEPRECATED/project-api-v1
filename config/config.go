package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	filename  = "./config.json"
	GetSample = `{
		"password_salt": "PASSWORD_SALT",
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
	URL          string `json:"frontend_url"`
	PasswordSalt string `json:"password_salt"`
	Database     struct {
		Url      string `json:"url"`
		Port     int    `json:"port"`
		DbName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
}

func Get() (*conf, error) {
	config, err := ioutil.ReadFile(filename)
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
