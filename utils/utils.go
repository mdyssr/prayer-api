package utils

import (
	"github.com/mdyssr/prayer-api/models"
	"math"
)

func GetNearestMethod(coords models.Coords, methods []models.PrayerMethod) int {

	//distances := make(map[int]float64)
	// methods is like: [{Id: 1, Lat: 10, Long: 20}, {Id: 0, Lat: 0, Long: 2}, ...]
	// distance is like >> {1: 1.14, 2: 517, ...}
	distances := make(map[int]float64)

	for _, m := range methods {
		latDiff := coords.Latitude - m.Coords.Latitude
		lonDiff := coords.Longitude - m.Coords.Longitude
		distance := math.Sqrt(math.Pow(latDiff, 2) + math.Pow(lonDiff, 2))
		distances[m.ID] = distance
	}

	// the default smallest Id is the first method
	closestMethodID := methods[0].ID
	smallestDistance := distances[closestMethodID]

	for id, distance := range distances {
		if smallestDistance > distance {
			smallestDistance = distance
			closestMethodID = id
		}
	}

	return closestMethodID
}