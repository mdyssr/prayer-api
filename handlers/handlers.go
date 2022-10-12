package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdyssr/prayer-api/models"
	"github.com/mdyssr/prayer-api/services"
	"github.com/mdyssr/prayer-api/utils"
)

const GetUserIPDataError = Error("Error getting user IP data")
const GetPrayerMethodsError = Error("Error getting prayer methods")
const GetPrayerDataError = Error("Error getting prayer data")

type Error string

func (e Error) Error() string {
	return string(e)
}

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"Welcome Home!"})
}

func GetPrayerTimes(c echo.Context) error {

	geoData, err := services.GetGeoLocation()
	if err != nil {
		return GetUserIPDataError
	}

	methods, err := services.GetMethods()
	if err != nil {
		return GetPrayerMethodsError
	}

	nearestMethodID := utils.GetNearestMethod(&geoData.Coords, methods)
	prayerTimesParams := &models.PrayerTimesParams{
		Coords:   geoData.Coords,
		MethodID: nearestMethodID,
	}

	prayerTimes, err := services.GetPrayerData(prayerTimesParams)
	if err != nil {
		return GetPrayerDataError
	}

	return c.JSON(http.StatusOK, prayerTimes)
}

// func GetPrayerTimes(c echo.Context) error {
// 	clientIP := c.RealIP()
// 	ipData, err := services.GetIPData(clientIP)
// 	if err != nil {
// 		return GetUserIPError
// 	}

// 	data, err := services.GetMethods()
// 	if err != nil {
// 		return GetPrayerMethodsError
// 	}

// 	nearestMethodID := utils.GetNearestMethod(ipData.Coords, data)
// 	//fmt.Println(nearestMethodID)
// 	prayerTimesParams := &models.PrayerTimesParams{
// 		Coords:   ipData.Coords,
// 		MethodID: nearestMethodID,
// 	}

// 	prayerTimes, err := services.GetPrayerData(prayerTimesParams)
// 	if err != nil {
// 		return GetPrayerDataError
// 	}

// 	return c.JSON(http.StatusOK, prayerTimes)
// }
