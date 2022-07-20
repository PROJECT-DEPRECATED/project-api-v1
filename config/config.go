package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	filename = "./config.json"
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

func Get() conf {
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var data conf
	err = json.Unmarshal(config, &data)
	if err != nil {
		panic(err)
	}

	return data
}
