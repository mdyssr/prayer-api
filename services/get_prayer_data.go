package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdyssr/prayer-api/data"
	"github.com/mdyssr/prayer-api/models"
	"io"
	"net/http"
	"os"
	"time"
)

const PRAYER_TIMINGS_BASE_URL = "https://api.aladhan.com/v1/timings/"

type PrayersResponse struct {
	Data Timings `json:"data"`
}

type Timings struct {
	Timings models.PrayerTimings `json:"timings"`
}

func GetPrayerData(params *models.PrayerTimesParams) (*models.PrayerTimings, error) {
	client := http.Client{}
	now := time.Now().Unix()
	url := fmt.Sprintf("%s%d?latitude=%g&longitude=%g&method=%d", PRAYER_TIMINGS_BASE_URL, now, params.Coords.Latitude, params.Coords.Longitude, params.MethodID)
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	prayersResponse := new(PrayersResponse)
	err = json.Unmarshal(bodyBytes, prayersResponse)
	if err != nil {
		return nil, err
	}

	return &prayersResponse.Data.Timings, nil
}

func GetMethods() ([]models.PrayerMethod, error) {
	//cwd, err := os.Getwd()
	//if err != nil {
	//	return nil, err
	//}
	//jsonFile := filepath.Join(cwd, "static", "methods_locations.json")
	//return getMethodsFromFile(jsonFile)
	//

	return data.PrayerMethods, nil
}

func getMethodsFromFile(file string) ([]models.PrayerMethod, error) {
	jsonFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var prayerMethods []models.PrayerMethod

	err = json.Unmarshal(jsonBytes, &prayerMethods)
	if err != nil {
		return nil, err
	}

	return prayerMethods, nil
}
