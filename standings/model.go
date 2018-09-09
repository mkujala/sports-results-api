package standings

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sports-results/db"

	"gopkg.in/mgo.v2/bson"
)

// Standings is representation of standings document in mongodb
type Standings struct {
	ID      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Team    string        `json:"team" bson:"team"`
	Season  int           `json:"season" bson:"season"`
	League  string        `json:"league" bson:"league"`
	Venue   string        `json:"venue" bson:"venue"`
	GP      int           `json:"gp" bson:"gp"`
	Wins    int           `json:"wins" bson:"wins"`
	Loses   int           `json:"loses" bson:"loses"`
	Draws   int           `json:"draws,omitempty" bson:"draws,omitempty"`
	OTLoses int           `json:"otLoses,omitempty" bson:"otLoses,omitempty"`
	OTWins  int           `json:"otWins,omitempty" bson:"otWins,omitempty"`
	GA      int           `json:"ga" bson:"ga"`
	GF      int           `json:"gf" bson:"gf"`
	PTS     int           `json:"pts" bson:"pts"`
}

// All reads full standings from DB (home & away games)
func allFromDB(league string, season int) ([]Standings, error) {
	stnds := []Standings{}
	err := db.Standings.Find(bson.M{"venue": "all", "league": league}).All(&stnds)
	if err != nil {
		return nil, err
	}
	return stnds, nil
}

// Home reads standings from DB for home games
func homeFromDB(league string, season int) ([]Standings, error) {
	stnds := []Standings{}
	err := db.Standings.Find(bson.M{"venue": "home"}).All(&stnds)
	if err != nil {
		return nil, err
	}
	return stnds, nil
}

// Away reads standings from DB for away games
func awayFromDB(league string, season int) ([]Standings, error) {
	stnds := []Standings{}
	err := db.Standings.Find(bson.M{"venue": "away"}).All(&stnds)
	if err != nil {
		return nil, err
	}
	return stnds, nil
}

// Insert standings to DB
func insertDB(r *http.Request) ([]Standings, error) {
	stnds := []Standings{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return stnds, err
	}

	err = json.Unmarshal(body, &stnds)
	if err != nil {
		return stnds, err
	}

	// convert []Standings{} to []interface{}
	s := make([]interface{}, len(stnds))
	for i, m := range stnds {
		s[i] = m
	}

	// insert documents
	err = db.Standings.Insert(s...)
	return stnds, err
}
