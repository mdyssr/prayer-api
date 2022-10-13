package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mdyssr/prayers"
	"net/http"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"Welcome Home!"})
}

func GetPrayerTimes(c echo.Context) error {
	prayerTimes, err := prayers.GetPrayersData()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, prayerTimes)
}
