package models

type Coords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoData struct {
	IP string `json:"ip"`
	Coords
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

type StandardPrayerTimeDesignation struct {
	Ar PrayerTimeLanguageDesignation
	En PrayerTimeLanguageDesignation
}
type PrayerTimeLanguageDesignation struct {
	Abbreviated string
	Expanded    string
}

type StandardPrayerTime struct {
	Value       string
	Designation StandardPrayerTimeDesignation
}

type PrayerTimeDetails struct {
	Military string
	Standard StandardPrayerTime
}

type FormattedPrayerTiming struct {
	Name string
	Time PrayerTimeDetails
}

type FormattedPrayerTimings []FormattedPrayerTiming

//
//type FormattedPrayerTimings struct {
//	Fajr    FormattedPrayerTiming `json:"fajr"`
//	Sunrise FormattedPrayerTiming `json:"sunrise"`
//	Dhuhr   FormattedPrayerTiming `json:"dhuhr"`
//	Asr     FormattedPrayerTiming `json:"asr"`
//	Sunset  FormattedPrayerTiming `json:"sunset"`
//	Maghrib FormattedPrayerTiming `json:"maghrib"`
//	Isha    FormattedPrayerTiming `json:"isha"`
//}

//type PrayerTimings struct {
//	Fajr    string `json:"Fajr"`
//	Sunrise string `json:"Sunrise"`
//	Dhuhr   string `json:"Dhuhr"`
//	Asr     string `json:"Asr"`
//	Sunset  string `json:"Sunset"`
//	Maghrib string `json:"Maghrib"`
//	Isha    string `json:"Isha"`
//}

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
	PrayerTimings FormattedPrayerTimings
	HijriDate     HijriDate
}

type PrayerTimesParams struct {
	Coords   Coords
	MethodID int
}
