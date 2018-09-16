package nhl

import (
	"net/http"
	"sports-results/formatter"
	"sports-results/leagues"
	"sports-results/standings"

	"github.com/julienschmidt/httprouter"
)

const league = "nhl"

// Standings for NHL
func Standings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stnds []standings.Standings
	venue := p.ByName("venue")
	season := p.ByName("season")
	conference := p.ByName("conference")
	stnds = leagues.StandingsFromDB(league, venue, season, conference)
	ls := addAverages(stnds)
	leagues.JSONout(w, ls)
}

// addAverages calculates NHL specific averages
func addAverages(s []standings.Standings) []Stats {
	var stats []Stats

	for _, j := range s {
		statLine := Stats{}
		statLine.Team = j.Team
		statLine.League = j.League
		statLine.Conference = j.Conference
		statLine.Season = j.Season
		statLine.Venue = j.Venue
		statLine.GP = j.GP
		statLine.Wins = j.Wins
		statLine.Loses = j.Loses
		statLine.OTLoses = j.OTLoses
		statLine.OTWins = j.OTWins
		statLine.StrWinPercent = formatter.Round2F(float64(j.Wins) / float64(j.GP))
		statLine.WinPercent = formatter.Round2F(float64(j.OTWins+j.Wins) / float64(j.GP))
		statLine.OTpercent = formatter.Round2F(float64(j.OTWins+j.OTLoses) / float64(j.GP))
		statLine.GA = j.GA
		statLine.GF = j.GF
		statLine.GAavg = formatter.Round2F((float64(j.GA) / float64(j.GP)))
		statLine.GFavg = formatter.Round2F((float64(j.GF) / float64(j.GP)))
		statLine.PTS = j.PTS
		statLine.PtsAvg = formatter.Round2F((float64(j.PTS) / float64(j.GP)))
		stats = append(stats, statLine)
	}

	return stats
}
