package korisliiga

import "gopkg.in/mgo.v2/bson"

// Stats is representation of NBA standings
type Stats struct {
	ID              bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Team            string        `json:"team" bson:"team"`
	Season          int           `json:"season" bson:"season"`
	League          string        `json:"league" bson:"league"`
	Venue           string        `json:"venue" bson:"venue"`
	GP              int           `json:"gp" bson:"gp"`
	Wins            int           `json:"wins" bson:"wins"`
	Loses           int           `json:"loses" bson:"loses"`
	WinPercent      float64       `json:"win_p,omitempty" bson:"win_p,omitempty"`
	PythagoreanWins float64       `json:"pythagoreanWins,omitempty" bson:"pythagoreanWins,omitempty"`
	GA              int           `json:"ga" bson:"ga"`
	GF              int           `json:"gf" bson:"gf"`
	GAavg           float64       `json:"gaAvg" bson:"gaAvg"`
	GFavg           float64       `json:"gfAvg" bson:"gfAvg"`
}
