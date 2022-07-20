package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/devproje/project-website/log"
	"github.com/gin-gonic/gin"
)

type hangangData struct {
	WaterPOS struct {
		Row []struct {
			Date   string `json:"MSR_DATE"`
			Time   string `json:"MSR_TIME"`
			SiteID string `json:"SITE_ID"`
			Temp   string `json:"W_TEMP"`
			PH     string `json:"W_PH"`
		} `json:"row"`
	} `json:"WPOSInformationTime"`
}

func getHangang(area string) (*hangangData, int) {
	var typeStr string
	switch area {
	case "tancheon":
		typeStr = "1"
	case "jungnangcheon":
		typeStr = "2"
	case "anyang":
		typeStr = "3"
	case "seonyu":
		typeStr = "4"
	case "noryangjin":
		typeStr = "5"
	}

	var url = "http://openapi.seoul.go.kr:8088/sample/json/WPOSInformationTime/" + typeStr + "/" + typeStr
	res, err := http.Get(url)
	if err != nil {
		log.Logger.Errorln(err)
		return nil, res.StatusCode
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger.Errorln(err)
		return nil, res.StatusCode
	}

	var data hangangData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Logger.Errorln(err)
		return nil, res.StatusCode
	}

	return &data, res.StatusCode
}

func Hangang(context *gin.Context) {
	area := context.Param("area")
	if area == "" {
		area = "jungnangcheon"
	}
	before := time.Now()
	hangang, status := getHangang(area)
	row := hangang.WaterPOS.Row[0]

	parse, err := time.Parse("20060102", row.Date)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	respondTime := time.Since(before)

	if status == 404 {
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"error":        http.StatusText(status),
		})
	}

	context.JSON(200, gin.H{
		"status":       status,
		"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
		"area":         row.SiteID,
		"date":         parse.Format("2006-01-02"),
		"time":         row.Time,
		"temp":         row.Temp,
		"ph":           row.PH,
	})
}
