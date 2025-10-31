package nascar_scraper

import (
	"testing"
)

// Test Get_race_data for wrong seasons and correct seasons
// for a valid return value.
func Test_get_data(t *testing.T) {
	GetRaceData(10, 900)
	GetRaceData(1999, 2005)
	GetRaceData(2024, 2026)
}
