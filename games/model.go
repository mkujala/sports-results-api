package games

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Games data model for game results and pre-game odds
type Games struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	League    string        `json:"league" bson:"league"`
	Season    int           `json:"season" bson:"season"`
	Date      time.Time     `json:"date" bson:"date"`
	HomeTeam  string        `json:"homeTeam" bson:"homeTeam"`
	AwayTeam  string        `json:"awayTeam" bson:"awayTeam"`
	HomeScore int           `json:"homeScore" bson:"homeScore"`
	AwayScore int           `json:"awayScore" bson:"awayScore"`
	OT        bool          `json:"ot,omitempty" bson:"ot,omitempty"`
	SO        bool          `json:"so,omitempty" bson:"so,omitempty"`
	HomeOdds  float64       `json:"homeOdds,omitempty" bson:"homeOdds,omitempty"`
	AwayOdds  float64       `json:"awayOdds,omitempty" bson:"awayOdds,omitempty"`
	DrawOdds  float64       `json:"drawOdds,omitempty" bson:"drawOdds,omitempty"`
	TotalLine float64       `json:"totalLine,omitempty" bson:"totalLine,omitempty"`
}
