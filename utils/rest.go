package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAPI[T any](url string) (*T, int) {
	res, err := http.Get(url)
	if err != nil {
		return nil, res.StatusCode
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode
	}

	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, res.StatusCode
	}

	return &data, res.StatusCode
}
