package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/utils"
	"github.com/gin-gonic/gin"
)

type HangangData struct {
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

func getHangang(area string) (*HangangData, int) {
	conf, _ := config.Get()
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
	default:
		return nil, 404
	}

	var url = fmt.Sprintf("%s/sample/json/WPOSInformationTime/%s/%s", conf.HangangAPI.URL, typeStr, typeStr)
	return utils.GetAPI[HangangData](url)
}

func Hangang(context *gin.Context) {
	area := context.Param("area")
	if area == "" {
		area = "jungnangcheon"
	}
	before := time.Now()
	hangang, status := getHangang(area)
	row := hangang.WaterPOS.Row[0]
	respondTime := time.Since(before)

	parse, err := time.Parse("20060102", row.Date)
	if err != nil {
		return
	}
	switch status {
	case 404:
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"error":        http.StatusText(status),
		})
	case 200:
		context.JSON(status, gin.H{
			"status":       status,
			"respond_time": strconv.FormatInt(respondTime.Milliseconds(), 10) + "ms",
			"area":         row.SiteID,
			"date":         parse.Format("2006-01-02"),
			"time":         row.Time,
			"temp":         row.Temp,
			"ph":           row.PH,
		})
	}
}
