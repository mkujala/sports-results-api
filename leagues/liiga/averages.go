package liiga

import (
	"net/http"
	"sports-results/formatter"
	"sports-results/leagues"
	"sports-results/standings"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type yearlyStnds struct {
	list map[string][]standings.Standings
}

type liigaAverages struct {
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

// Averages returns liiga specific averages for selected venue and seasons
func Averages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	venue := p.ByName("venue")
	seasons := strings.Split(p.ByName("seasons"), ",")
	stnds := getSeasons(venue, seasons)

	// averages for top2
	averagesTeams(stnds, 2)

	// averages for top6
	averagesTeams(stnds, 6)

	// averages for 7-10

	// averages for bottom3
}

// GetSeasons get n last seasons sorted by points
func getSeasons(venue string, seasons []string) yearlyStnds {
	var stnds = yearlyStnds{list: make(map[string][]standings.Standings)}

	for _, season := range seasons {
		stnds.list[season] = leagues.StandingsFromDB(league, venue, season, conference)
	}
	return stnds
}

func averagesTeams(s yearlyStnds, count int) {
	for _, team := range s.list {
		seasonAverages(team, count)
	}
}

// wins, loses, otwins, otloses, win%reg, win%all, ot%, gf, ga, gf/gp, ga/gp, pts, pts/gp
func seasonAverages(s []standings.Standings, count int) {
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

	for i := 0; i < count; i++ {
		wins = append(wins, float64(s[i].Wins))
		loses = append(loses, float64(s[i].Loses))
		otwins = append(otwins, float64(s[i].OTWins))
		otloses = append(otloses, float64(s[i].OTLoses))
		winRegP = append(winRegP, float64(s[i].Wins)/float64(s[i].GP))
		winAllP = append(winAllP, (float64(s[i].Wins)+float64(s[i].OTWins))/float64(s[i].GP))
		otP = append(otP, (float64(s[i].OTLoses)+float64(s[i].OTWins))/float64(s[i].GP))
		gf = append(gf, float64(s[i].GF))
		ga = append(ga, float64(s[i].GA))
		gfGP = append(gfGP, float64(s[i].GF)/float64(s[i].GP))
		gaGP = append(gaGP, float64(s[i].GA)/float64(s[i].GP))
		points = append(points, float64(s[i].PTS))
		ptsGP = append(ptsGP, float64(s[i].PTS)/float64(s[i].GP))
	}

	var output liigaAverages
	output.Season = s[0].Season
	output.Description = "top " + strconv.Itoa(count)
	output.Wins = average(wins...)
	output.Loses = average(loses...)
	output.Otwins = average(otwins...)
	output.Otloses = average(otloses...)
	output.WinRegP = average(winRegP...)
	output.WinAllP = average(winAllP...)
	output.OtP = average(otP...)
	output.GF = average(gf...)
	output.GA = average(ga...)
	output.GFgp = average(gfGP...)
	output.GAgp = average(gaGP...)
	output.Points = average(points...)
	output.PTSgp = average(ptsGP...)

	_ = formatter.PrettyPrint(output)
}

// average for any number of arguments
func average(nums ...float64) float64 {
	divider := float64(len(nums))
	var total float64
	for _, val := range nums {
		total = total + val
	}
	return formatter.Round2F(total / divider)
}
