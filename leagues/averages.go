package leagues

import (
	"sports-results/formatter"
	"sports-results/standings"
)

// YearlyStnds is used to create map with multiple seasons standings
type YearlyStnds struct {
	list map[string][]standings.Standings
}

// GetSeasons get n last seasons sorted by points
func GetSeasons(venue string, seasons []string) YearlyStnds { // modify liiga to use this!!!!!!!!!!!!
	var stnds = YearlyStnds{list: make(map[string][]standings.Standings)}
	for _, season := range seasons {
		stnds.list[season] = StandingsFromDB(league, venue, season, conference) // nyt const liiga/standings etc
	}
	return stnds
}

// Average for any number of arguments
func Average(nums ...float64) float64 {
	divider := float64(len(nums))
	var total float64
	for _, val := range nums {
		total = total + val
	}
	return formatter.Round2F(total / divider)
}
