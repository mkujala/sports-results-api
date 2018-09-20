package liiga

import (
	"fmt"
	"net/http"
	"sports-results/leagues"
	"sports-results/standings"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func Averages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	venue := p.ByName("venue")
	seasons := strings.Split(p.ByName("seasons"), ",")
	getSeasons(venue, seasons)
}

// GetSeasons get n last seasons sorted by points
func getSeasons(venue string, seasons []string) {
	stnds := make(map[string][]standings.Standings)

	for _, season := range seasons {
		stnds[season] = leagues.StandingsFromDB(league, venue, season, conference)
	}
	fmt.Println(stnds)
}

func top2Averages() {
	// WIP
	// calc averages for top2 teams for all, home, away
	// wins, loses, otwins, otloses, gf, ga, pts, pts/gp, gf/gp, ga/gp, win%, ot%
}

// averages for top6

// averages for 7-10

// averages for bottom3
