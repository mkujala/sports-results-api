package liiga

import "gopkg.in/mgo.v2/bson"

// Stats is representation of standings document in mongodb
type Stats struct {
	ID      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Team    string        `json:"team" bson:"team"`
	Season  int           `json:"season" bson:"season"`
	League  string        `json:"league" bson:"league"`
	Venue   string        `json:"venue" bson:"venue"`
	GP      int           `json:"gp" bson:"gp"`
	Wins    int           `json:"wins" bson:"wins"`
	Loses   int           `json:"loses" bson:"loses"`
	OTLoses int           `json:"otLoses,omitempty" bson:"otLoses,omitempty"`
	OTWins  int           `json:"otWins,omitempty" bson:"otWins,omitempty"`
	GA      int           `json:"ga" bson:"ga"`
	GF      int           `json:"gf" bson:"gf"`
	PTS     int           `json:"pts" bson:"pts"`
	PtsAvg  float32       `json:"ptsAvg,omitempty" bson:"ptsAvg,omitempty"`
}
