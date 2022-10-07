package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mdyssr/prayer-api/models"
	"github.com/mdyssr/prayer-api/services"
	"github.com/mdyssr/prayer-api/utils"
	"net/http"
)

const GetUserIPError = Error("Error getting user IP")
const GetPrayerMethodsError = Error("Error getting prayer methods")
const GetPrayerTimesError = Error("Error getting prayer times")

type Error string

func (e Error) Error() string {
	return string(e)
}

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"Welcome Home!"})
}

func GetPrayerTimes(e echo.Context) error {
	ipData, err := services.GetIPData()
	if err != nil {
		return GetUserIPError
	}

	data, err := services.GetMethods()
	if err != nil {
		return GetPrayerMethodsError
	}

	// key: method, value: location:
	nearestMethodID := utils.GetNearestMethod(ipData.Coords, data)
	//fmt.Println(nearestMethodID)
	prayerTimesParams := &models.PrayerTimesParams{
		Coords:   ipData.Coords,
		MethodID: nearestMethodID,
	}

	prayerTimes, err := services.GetPrayerData(prayerTimesParams)
	if err != nil {
		return GetPrayerTimesError
	}

	return e.JSON(http.StatusOK, prayerTimes)
}

/*

	ipData, err := services.GetIPData()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := services.GetMethods()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// key: method, value: location:
	nearestMethodID := utils.GetNearestMethod(ipData.Coords, data)
	//fmt.Println(nearestMethodID)
	prayerTimesParams := &models.PrayerTimesParams{
		Coords:   ipData.Coords,
		MethodID: nearestMethodID,
	}

	prayerTimes, err := services.GetPrayerData(prayerTimesParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(prayerTimes)

*/
