package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/utils"
	"github.com/gin-gonic/gin"
)

type uuidExtractor struct {
	Username string `json:"name"`
	UniqueId string `json:"id"`
}

type EncodedMojangAPI struct {
	Username   string `json:"name"`
	UniqueId   string `json:"id"`
	Properties []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
}

type MojangAPI struct {
	UniqueId string `json:"profileId"`
	Username string `json:"profileName"`
	Textures struct {
		Skin struct {
			Url string `json:"url"`
		} `json:"SKIN"`
	} `json:"textures"`
}

func getUniqueId(username string) (*uuidExtractor, int) {
	conf, _ := config.Get()
	url := fmt.Sprintf("%s/users/profiles/minecraft/%s", conf.MojangAPI.API, username)

	return utils.GetAPI[uuidExtractor](url)
}

func decodeProperty(encoded EncodedMojangAPI) (*MojangAPI, error) {
	bytes, err := base64.StdEncoding.DecodeString(encoded.Properties[0].Value)
	if err != nil {
		return nil, err
	}

	var decoded MojangAPI
	err = json.Unmarshal(bytes, &decoded)
	if err != nil {
		return nil, err
	}

	return &decoded, nil
}

func getMCProfile(username string) (*MojangAPI, int) {
	conf, _ := config.Get()
	var uniqueId, status = getUniqueId(username)
	if status != 200 {
		return nil, status
	}

	url := fmt.Sprintf("%s/session/minecraft/profile/%s", conf.MojangAPI.Session, uniqueId.UniqueId)

	e, status := utils.GetAPI[EncodedMojangAPI](url)
	data, err := decodeProperty(*e)
	if err != nil {
		return nil, 500
	}

	return data, status
}

func MCProfile(context *gin.Context) {
	username := context.Param("username")
	before := time.Now()
	mojang, status := getMCProfile(username)
	respondTime := time.Since(before)

	switch status {
	case 404:
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"error":        http.StatusText(status),
		})
	case 400:
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"error":        http.StatusText(status),
		})
	case 204:
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"error":        http.StatusText(status),
		})
	case 200:
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"username":     mojang.Username,
			"unique_id":    mojang.UniqueId,
			"skin_url":     mojang.Textures.Skin.Url,
		})
	}
}
