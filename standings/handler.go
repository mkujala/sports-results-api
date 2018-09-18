package standings

import (
	"encoding/json"
	"fmt"
	"net/http"

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
