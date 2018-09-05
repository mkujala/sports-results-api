package games

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Games struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	League    string        `json:"league" bson:"league"`
	HomeTeam  string        `json:"homeTeam" bson:"homeTeam"`
	AwayTeam  string        `json:"awayTeam" bson:"awayTeam"`
	HomeScore int           `json:"homeScore" bson:"homeScore"`
	AwayScore int           `json:"awayScore" bson:"awayScore"`
	Date      time.Time     `json:"date" bson:"date"`
	OT        bool          `json:"ot" bson:"ot"`
	SO        bool          `json:"so" bson:"so"`
}
