package epl

import (
	"net/http"
	"sports-results/formatter"
	"sports-results/leagues"
	"sports-results/standings"

	"github.com/julienschmidt/httprouter"
)

// Standings for EPL
func Standings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stnds []standings.Standings
	venue := p.ByName("venue")
	season := p.ByName("season")
	stnds = leagues.StandingsFromDB("epl", venue, season)

	ls := addAverages(stnds)
	leagues.JsonOut(w, ls)
}

// addAverages calculates liiga specific averages
func addAverages(s []standings.Standings) []Stats {
	var stats []Stats

	for _, j := range s {
		statLine := Stats{}
		statLine.Team = j.Team
		statLine.League = j.League
		statLine.Season = j.Season
		statLine.Venue = j.Venue
		statLine.GP = j.GP
		statLine.Wins = j.Wins
		statLine.Loses = j.Loses
		statLine.Draws = j.Draws
		statLine.WinPercent = formatter.Round2F(float64(j.Wins) / float64(j.GP))
		statLine.DrawPpercent = formatter.Round2F(float64(j.Draws) / float64(j.GP))
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
