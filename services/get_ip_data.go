package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdyssr/prayer-api/models"
	"io"
	"net/http"
)

//const IP_URL = "https://ipapi.co/"

// type IPData struct {
// 	IP string `json:"ip"`
// 	models.Coords
// }
//
//func GetIPData(ip string) (*IPData, error) {
//	url := fmt.Sprintf("%s%s/json", IP_URL, ip)
//	client := http.Client{}
//	request, err := http.NewRequest("GET", url, nil)
//	request.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
//
//	response, err := client.Do(request)
//	defer response.Body.Close()
//
//	if err != nil {
//		return nil, err
//	}
//	responseData, err := io.ReadAll(response.Body)
//	if err != nil {
//		return nil, err
//	}
//	ipData := new(IPData)
//	err = json.Unmarshal(responseData, ipData)
//	if err != nil {
//		return nil, err
//
//	}
//
//	return ipData, nil
//}

// GetGeoLocation should return Lat, Long values for the current requester

func GetGeoLocation() (*models.GeoData, error) {
	const base_url = "http://api.ipstack.com/check"
	const api_key = "10c6be7bf5f2704393a672e2c714ca04"
	const fields = "ip,latitude,longitude"
	//const url = "http://api.ipstack.com/check?access_key=10c6be7bf5f2704393a672e2c714ca04&fields=ip,latitude,longitude"
	url := fmt.Sprintf("%s?access_key=%s&fields=%s", base_url, api_key, fields)
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	geoData := new(models.GeoData)
	err = json.Unmarshal(responseData, geoData)
	if err != nil {
		return nil, err
	}
	fmt.Println(geoData.IP)
	return geoData, nil
}
