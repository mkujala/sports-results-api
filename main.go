package main

import (
	"fmt"
	"log"
	"net/http"
	"sports-results/leagues/epl"
	"sports-results/leagues/liiga"
	"sports-results/standings"

	"github.com/julienschmidt/httprouter"
)

// Constants
const (
	apiURL = "/api/v1"
	port   = ":8000"
)

func main() {
	router := httprouter.New()

	router.GET(apiURL+"/", hello)
	router.GET(apiURL+"/standings/:venue", standings.Get)
	router.POST(apiURL+"/standings", standings.Insert)
	router.GET(apiURL+"/liiga/standings/:season/:venue", liiga.Standings)
	router.GET(apiURL+"/epl/standings/:season/:venue", epl.Standings)
	// router.DELETE(apiURL+"/standings/:id", standings.Delete)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Server start failed when using PORT:", port)
	}
}

func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello!\n")
}
