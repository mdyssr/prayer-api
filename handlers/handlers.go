package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdyssr/prayer-api/models"
	"github.com/mdyssr/prayer-api/services"
	"github.com/mdyssr/prayer-api/utils"
)

const NoIPProvidedError = Error("No IP Provided")
const InvalidIPError = Error("Invalid IP Address")
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
	clientIP := c.QueryParam("ip")
	if clientIP == "" {
		return NoIPProvidedError
	}
	ip := net.ParseIP(clientIP)
	if ip == nil {
		return InvalidIPError
	}
	ipData, err := services.GetIPData(ip.String())
	if err != nil {
		return GetUserIPDataError
	}

	if ipData.IPDataError.Error {
		return GetUserIPDataError
	}

	fmt.Println(ip.String())

	data, err := services.GetMethods()
	if err != nil {
		return GetPrayerMethodsError
	}

	nearestMethodID := utils.GetNearestMethod(ipData.Coords, data)
	prayerTimesParams := &models.PrayerTimesParams{
		Coords:   ipData.Coords,
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
