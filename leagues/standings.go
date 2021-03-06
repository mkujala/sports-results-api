package leagues

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sports-results/standings"
	"strconv"
)

type leagueData interface{}

// StandingsFromDB generalizes data fetch from DB for all leagues
func StandingsFromDB(league, venue, season, conference string) []standings.Standings {
	var stnds []standings.Standings
	iSeason, err := strconv.Atoi(season)
	if err != nil {
		log.Fatal("error: invalid season, " + err.Error())
		return nil
	}

	switch venue {
	case "all", "home", "away":
		stnds, err = standings.GetFromDB(league, venue, iSeason, conference)
	default:
		log.Fatal("error: invalid venue, " + err.Error())
		return nil
	}
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return stnds
}

// JSONout marshals data struct and outputs json
func JSONout(w http.ResponseWriter, data leagueData) {
	sjson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", sjson)
}
