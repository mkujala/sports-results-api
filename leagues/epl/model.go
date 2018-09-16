package epl

import "gopkg.in/mgo.v2/bson"

// Stats is representation of standings document in mongodb
type Stats struct {
	ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Team         string        `json:"team" bson:"team"`
	Season       int           `json:"season" bson:"season"`
	League       string        `json:"league" bson:"league"`
	Venue        string        `json:"venue" bson:"venue"`
	GP           int           `json:"gp" bson:"gp"`
	Wins         int           `json:"wins" bson:"wins"`
	Loses        int           `json:"loses" bson:"loses"`
	Draws        int           `json:"draws" bson:"draws"`
	WinPercent   float64       `json:"win_p,omitempty" bson:"win_p,omitempty"`
	DrawPpercent float64       `json:"draw_p,omitempty" bson:"draw_p,omitempty"`
	GA           int           `json:"ga" bson:"ga"`
	GF           int           `json:"gf" bson:"gf"`
	GAavg        float64       `json:"gaAvg" bson:"gaAvg"`
	GFavg        float64       `json:"gfAvg" bson:"gfAvg"`
	PTS          int           `json:"pts" bson:"pts"`
	PtsAvg       float64       `json:"ptsAvg,omitempty" bson:"ptsAvg,omitempty"`
}
