package liiga

import (
	"fmt"
	"net/http"
	"sports-results/leagues"
	"sports-results/standings"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type yearlyStnds struct {
	list map[string][]standings.Standings
}

func Averages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	venue := p.ByName("venue")
	seasons := strings.Split(p.ByName("seasons"), ",")
	getSeasons(venue, seasons)
}

// GetSeasons get n last seasons sorted by points
func getSeasons(venue string, seasons []string) {
	var stnds = yearlyStnds{list: make(map[string][]standings.Standings)}

	for _, season := range seasons {
		stnds.list[season] = leagues.StandingsFromDB(league, venue, season, conference)
	}
	// formatter.PrettyPrint(stnds.list)
	top2Averages(stnds)
}

func top2Averages(s yearlyStnds) {
	// WIP
	// calc averages for top2 teams for all, home, away
	// wins, loses, otwins, otloses, gf, ga, pts, pts/gp, gf/gp, ga/gp, win%, ot%
	for _, season := range s.list {
		fmt.Println(season[0].Team, season[1].Team, "pts avg top2", (float64(season[0].PTS)+float64(season[1].PTS))/2)
	}
}

// averages for top6

// averages for 7-10

// averages for bottom3
