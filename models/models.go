package models

type Coords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PrayerMethod struct {
	ID int `json:"id"`
	Coords
}

type PrayerTimings struct {
	Fajr    string `json:"Fajr"`
	Sunrise string `json:"Sunrise"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Sunset  string `json:"Sunset"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type HijriDate struct {
	Day     string `json:"day"`
	Weekday struct {
		En string `json:"en"`
		Ar string `json:"ar"`
	} `json:"weekday"`
	Month struct {
		Number int    `json:"number"`
		En     string `json:"en"`
		Ar     string `json:"ar"`
	} `json:"month"`
	Year string `json:"year"`
}

type PrayersData struct {
	PrayerTimings PrayerTimings
	HijriDate     HijriDate
}

type PrayerTimesParams struct {
	Coords   Coords
	MethodID int
}
