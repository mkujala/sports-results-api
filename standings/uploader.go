package standings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Insert new standings to DB
func Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Body == nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	stnds, err := insertDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sjson, err := json.Marshal(stnds)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", sjson)
}

// Replace can be used to insert standings during season (remove existing and add all as new)
func Replace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Body == nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	league := p.ByName("league")
	venue := p.ByName("venue")
	season := p.ByName("season")
	conference := p.ByName("conference")
	iSeason, err := strconv.Atoi(season)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// remove from DB, conference NULL = no conferences in league
	num, err := removeFromDB(league, venue, conference, iSeason)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if num < 1 {
		http.Error(w, "no documents found to replace", http.StatusInternalServerError)
		return
	}
	fmt.Printf("removed %d documents from %s, %d:\n", num, league, iSeason)

	// After removing from DB insert new standings
	stnds, err := insertDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sjson, err := json.Marshal(stnds)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", sjson)
}
