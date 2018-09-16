package leagues

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sports-results/standings"
	"strconv"
)

type Data interface{}

// Standings generalizes
func StandingsFromDB(league, venue, season string) []standings.Standings {
	var stnds []standings.Standings
	iSeason, err := strconv.Atoi(season)
	if err != nil {
		log.Fatal("error: invalid season, " + err.Error())
		return nil
	}

	switch venue {
	case "all", "home", "away":
		stnds, err = standings.GetFromDB(league, venue, iSeason)
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

func JsonOut(w http.ResponseWriter, data Data) {

	sjson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", sjson)

}
