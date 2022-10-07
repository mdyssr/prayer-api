package utils

import (
	"github.com/mdyssr/prayer-api/models"
	"testing"
)

func TestGetNearestMethod(t *testing.T) {
	currentCoords := models.Coords{Latitude: 40, Longitude: 60}
	coords := []models.PrayerMethod{
		{
			ID: 1,
			Coords: models.Coords{
				Latitude:  100,
				Longitude: 0,
			},
		},
		{
			ID: 0,
			Coords: models.Coords{
				Latitude:  50,
				Longitude: 30,
			},
		},
		{
			ID: 1,
			Coords: models.Coords{
				Latitude:  0,
				Longitude: 2,
			},
		},
	}

	got := GetNearestMethod(currentCoords, coords)
	want := coords[1].ID

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
