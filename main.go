package main

import (
	"log"
	"net/http"
	"sports-results/config"
	"sports-results/leagues/epl"
	"sports-results/leagues/korisliiga"
	"sports-results/leagues/liiga"
	"sports-results/leagues/nba"
	"sports-results/leagues/nhl"
	"sports-results/standings"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := config.Values()
	apiURL := config.APIURL
	port := config.Port

	router := httprouter.New()
	router.POST(apiURL+"/standings", standings.Insert)
	router.POST(apiURL+"/standings/replace/:league/:venue/:season/:conference", standings.Replace) // use conference value NULL, when league doesn't have conferences
	router.GET(apiURL+"/liiga/standings/:venue/:season", liiga.Standings)
	router.GET(apiURL+"/liiga/averages/:venue/:seasons", liiga.Averages) // seasons -> 20162017,20172018,...
	router.GET(apiURL+"/epl/standings/:venue/:season", epl.Standings)
	router.GET(apiURL+"/nhl/standings/:venue/:season/:conference", nhl.Standings)
	router.GET(apiURL+"/nhl/averages/:venue/:seasons/:conference", nhl.Averages) // seasons -> 20162017,20172018,...
	router.GET(apiURL+"/nba/standings/:venue/:season/:conference", nba.Standings)
	router.GET(apiURL+"/koris/standings/:venue/:season", korisliiga.Standings)
	router.GET(apiURL+"/koris/averages/:venue/:seasons", korisliiga.Averages) // seasons -> 20162017,20172018,...

	// router.DELETE(apiURL+"/standings/:id", standings.Delete)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Server start failed when using PORT:", port)
	}
}
