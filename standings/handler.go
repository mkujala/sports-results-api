package standings

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAll standings from DB
func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stnds, err := allFromDB()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	sjson, err := json.Marshal(stnds)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", sjson)
}

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
