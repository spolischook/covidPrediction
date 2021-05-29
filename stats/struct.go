package stats

import "time"

type CovidStats struct {
	CountryIso string
	Date       time.Time
	NewCases   int
}
