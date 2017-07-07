package models

import "time"

// TimeZone represents a single time zone object
type TimeZone struct {
	CountryIso2Code string        // CountryCode
	TimeZoneID      string        // TimeZoneId
	GmtOffset       time.Duration // GMT offset 1. Jan 2016
	DstOffset       time.Duration // DST offset 1. Jul 2016
	RawOffset       time.Duration // rawOffset (independant of DST)
}
