package korisliiga

import (
	"math"
	"net/http"
	"sports-results/formatter"
	"sports-results/leagues"
	"sports-results/standings"

	"github.com/julienschmidt/httprouter"
)

// Standings for Korisliiga
func Standings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stnds []standings.Standings
	venue := p.ByName("venue")
	season := p.ByName("season")
	stnds = leagues.StandingsFromDB(league, venue, season, conference)
	ls := addAverages(stnds)
	leagues.JSONout(w, ls)
}

// addAverages calculates Korisliiga specific averages
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
		statLine.WinPercent = formatter.Round2F(float64(j.OTWins+j.Wins) / float64(j.GP))
		statLine.PythagoreanWins = formatter.Round2F(pythagWin(j.GF, j.GA) * float64(j.GP))
		statLine.GA = j.GA
		statLine.GF = j.GF
		statLine.GAavg = formatter.Round2F((float64(j.GA) / float64(j.GP)))
		statLine.GFavg = formatter.Round2F((float64(j.GF) / float64(j.GP)))
		stats = append(stats, statLine)
	}

	return stats
}

// pyhthagWin calculates Pythagorean win%
// http://wikipedia.org/wiki/Pythagorean_expectation
func pythagWin(gf, ga int) float64 {
	gaExp := math.Pow(float64(ga), pythExp)
	gfExp := math.Pow(float64(gf), pythExp)
	return gfExp / (gfExp + gaExp)
}
