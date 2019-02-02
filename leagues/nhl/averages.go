package nhl

import (
	"net/http"
	"sports-results/leagues"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// wins, loses, otwins, otloses, win%reg, win%all, ot%, gf, ga, gf/gp, ga/gp, pts, pts/gp
type nhlAverages struct {
	Season      int
	Description string
	Wins,
	Loses,
	Otwins,
	Otloses,
	WinRegP,
	WinAllP,
	OtP,
	GF,
	GA,
	GFgp,
	GAgp,
	Points,
	PTSgp float64
}

// Averages returns league specific averages for selected venue and seasons
func Averages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	venue := p.ByName("venue")
	conference := p.ByName("conference")
	seasons := strings.Split(p.ByName("seasons"), ",")
	stnds := leagues.GetSeasons(league, venue, conference, seasons)

	// averages for top2
	response := seasonAverages(stnds, 0, 2, "top2_"+conference)

	// averages for top6
	response = append(response, seasonAverages(stnds, 0, 6, "top6_"+conference)...)

	// averages for 7-10
	response = append(response, seasonAverages(stnds, 6, 10, "top7to10_"+conference)...)

	// averages for bottom3
	response = append(response, seasonAverages(stnds, 3, 0, "bottom3_"+conference)...)

	leagues.JSONout(w, response)
}

// seasonAverages takes in YearlyStandings, start, end int, name string
// it returns slice of calculated nhlAverages
// if end != 0 -> start is index of first team and end is index of last team to count in averages
// if end == 0 -> start is count for how many teams from the end to count in averages
// 		example: end = 0, start = 3 -> count averages for bottom 3 teams
// name is used as Description in nhlAverages
func seasonAverages(s leagues.YearlyStnds, start, end int, name string) []nhlAverages {
	sAvg := []nhlAverages{}
	for _, team := range s.List {
		var (
			wins,
			loses,
			otwins,
			otloses,
			winRegP,
			winAllP,
			otP,
			gf,
			ga,
			gfGP,
			gaGP,
			points,
			ptsGP []float64
		)
		if end == 0 { // TRUE -> select teams from bottom and start tells how many
			end = len(team)
			start = end - start
		}

		for i := start; i < end; i++ {
			wins = append(wins, float64(team[i].Wins))
			loses = append(loses, float64(team[i].Loses))
			otwins = append(otwins, float64(team[i].OTWins))
			otloses = append(otloses, float64(team[i].OTLoses))
			winRegP = append(winRegP, float64(team[i].Wins)/float64(team[i].GP))
			winAllP = append(winAllP, (float64(team[i].Wins)+float64(team[i].OTWins))/float64(team[i].GP))
			otP = append(otP, (float64(team[i].OTLoses)+float64(team[i].OTWins))/float64(team[i].GP))
			gf = append(gf, float64(team[i].GF))
			ga = append(ga, float64(team[i].GA))
			gfGP = append(gfGP, float64(team[i].GF)/float64(team[i].GP))
			gaGP = append(gaGP, float64(team[i].GA)/float64(team[i].GP))
			points = append(points, float64(team[i].PTS))
			ptsGP = append(ptsGP, float64(team[i].PTS)/float64(team[i].GP))
		}

		var output nhlAverages
		output.Season = team[0].Season
		output.Description = name
		output.Wins = leagues.Average(wins...)
		output.Loses = leagues.Average(loses...)
		output.Otwins = leagues.Average(otwins...)
		output.Otloses = leagues.Average(otloses...)
		output.WinRegP = leagues.Average(winRegP...)
		output.WinAllP = leagues.Average(winAllP...)
		output.OtP = leagues.Average(otP...)
		output.GF = leagues.Average(gf...)
		output.GA = leagues.Average(ga...)
		output.GFgp = leagues.Average(gfGP...)
		output.GAgp = leagues.Average(gaGP...)
		output.Points = leagues.Average(points...)
		output.PTSgp = leagues.Average(ptsGP...)

		sAvg = append(sAvg, output)
	}
	return sAvg
}
