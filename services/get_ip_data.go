package services

import (
	"encoding/json"
	"github.com/mdyssr/prayer-api/models"
	"io"
	"net/http"
)

const IP_URL = "https://ipapi.co/json/"

type IPData struct {
	IP string `json:"ip"`
	models.Coords
}

func GetIPData() (*IPData, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", IP_URL, nil)
	request.Header.Set("User-Agent", "ipapi.co/#go-v1.5")

	response, err := client.Do(request)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	ipData := new(IPData)
	err = json.Unmarshal(responseData, ipData)
	if err != nil {
		return nil, err

	}
	return ipData, nil
}
