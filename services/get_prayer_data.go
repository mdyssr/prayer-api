package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdyssr/prayer-api/data"
	"github.com/mdyssr/prayer-api/models"
	"io"
	"net/http"
	"time"
)

const PRAYER_TIMINGS_BASE_URL = "https://api.aladhan.com/v1/timings/"

type PrayersResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Date    Date                 `json:"date"`
	Timings models.PrayerTimings `json:"timings"`
}

type Date struct {
	HijriDate models.HijriDate `json:"hijri"`
}

func GetPrayerData(params *models.PrayerTimesParams) (*models.PrayersData, error) {
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

	prayerData := models.PrayersData{
		HijriDate:     prayersResponse.Data.Date.HijriDate,
		PrayerTimings: prayersResponse.Data.Timings,
	}
	return &prayerData, nil
}

func GetMethods() ([]models.PrayerMethod, error) {
	return data.PrayerMethods, nil
}
