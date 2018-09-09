package liiga

import (
	"fmt"
	"net/http"
	"sports-results/standings"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Standings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stnds []standings.Standings
	venue := p.ByName("venue")
	season, err := strconv.Atoi(p.ByName("season"))
	if err != nil {
		http.Error(w, "error: invalid season, "+err.Error(), http.StatusInternalServerError)
		return
	}

	switch p.ByName("venue") {
	case "all", "home", "away":
		stnds, err = standings.GetFromDB("liiga", venue, season)
	default:
		http.Error(w, "error: invalid venue", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%#v\n", stnds)

	// testing pts avg
	fmt.Fprintf(w, "\n%v, ptsAvg: ", stnds[2].Team)
	fmt.Fprintf(w, "%.2f\n", float32(stnds[2].PTS)/float32(stnds[2].GP))
}

// PtsAvg calculates pts average per game
func PtsAvg(s []standings.Standings) {
	//-----------------
	// WORK IN PROGRESS
	//-----------------
	/*
		stats := []Stats{}
		for i, j := range s {
			stats[i] = s[i]
		}
	*/
}
