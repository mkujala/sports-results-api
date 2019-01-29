package nhl

import "gopkg.in/mgo.v2/bson"

// Stats is representation of NHL standings
type Stats struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Team          string        `json:"team" bson:"team"`
	Season        int           `json:"season" bson:"season"`
	League        string        `json:"league" bson:"league"`
	Conference    string        `json:"conference" bson:"conference"`
	Venue         string        `json:"venue" bson:"venue"`
	GP            int           `json:"gp" bson:"gp"`
	Wins          int           `json:"wins" bson:"wins"`
	Loses         int           `json:"loses" bson:"loses"`
	OTLoses       int           `json:"otLoses" bson:"otLoses"`
	OTWins        int           `json:"otWins" bson:"otWins"`
	StrWinPercent float64       `json:"strWin_p" bson:"strWin_p"`
	WinPercent    float64       `json:"win_p" bson:"win_p"`
	OTpercent     float64       `json:"ot_p" bson:"ot_p"`
	GA            int           `json:"ga" bson:"ga"`
	GF            int           `json:"gf" bson:"gf"`
	GAavg         float64       `json:"gaAvg" bson:"gaAvg"`
	GFavg         float64       `json:"gfAvg" bson:"gfAvg"`
	PTS           int           `json:"pts" bson:"pts"`
	PtsAvg        float64       `json:"ptsAvg" bson:"ptsAvg"`
}
