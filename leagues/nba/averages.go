package nba

import (
	"net/http"
	"sports-results/leagues"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//-----------------
// WORK IN PROGRESS
//-----------------

// nbaAverages struct copied from korisliiga == check data fields
//-----------------

// wins, loses, otwins, otloses, win%reg, win%all, ot%, gf, ga, gf/gp, ga/gp, pts, pts/gp
type nbaAverages struct {
	Season      int
	Description string
	Wins,
	Loses,
	WinAllP,
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
	seasons := strings.Split(p.ByName("seasons"), ",")
	conference := p.ByName("conference")
	stnds := leagues.GetSeasons(league, venue, conference, seasons)

	// averages for top2
	response := seasonAverages(stnds, 0, 2, "top2")

	// averages for top4
	response = append(response, seasonAverages(stnds, 0, 4, "top6")...)

	// averages for 5-8
	response = append(response, seasonAverages(stnds, 4, 8, "top5to8")...)

	// averages for bottom2
	response = append(response, seasonAverages(stnds, 2, 0, "bottom2")...)

	leagues.JSONout(w, response)
}

// seasonAverages takes in yearlyStandings, start, end int, name string
// it returns slice of calculated nbaAverages
// if end != 0 -> start is index of first team and end is index of last team to count in averages
// if end == 0 -> start is count for how many teams from the end to count in averages
// 		example: end = 0, start = 3 -> count averages for bottom 3 teams
// name is used as Description in nbaAverages
func seasonAverages(s leagues.YearlyStnds, start, end int, name string) []nbaAverages {
	sAvg := []nbaAverages{}
	for _, team := range s.List {
		var (
			wins,
			loses,
			winAllP,
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
			winAllP = append(winAllP, (float64(team[i].Wins) / float64(team[i].GP)))
			gf = append(gf, float64(team[i].GF))
			ga = append(ga, float64(team[i].GA))
			gfGP = append(gfGP, float64(team[i].GF)/float64(team[i].GP))
			gaGP = append(gaGP, float64(team[i].GA)/float64(team[i].GP))
			points = append(points, float64(team[i].PTS))
			ptsGP = append(ptsGP, float64(team[i].PTS)/float64(team[i].GP))
		}

		var output nbaAverages
		output.Season = team[0].Season
		output.Description = name
		output.Wins = leagues.Average(wins...)
		output.Loses = leagues.Average(loses...)
		output.WinAllP = leagues.Average(winAllP...)
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
