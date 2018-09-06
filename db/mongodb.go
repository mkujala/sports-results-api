package db

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// DB session
var DB *mgo.Database

// Games Collection
var Games *mgo.Collection

// Standings Collection
var Standings *mgo.Collection

func init() {
	const (
		mongoURI = "mongodb://localhost"
		mongoDB  = "sport-results"
	)

	s, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB(mongoDB)
	Games = DB.C("games")
	Standings = DB.C("standings")

	fmt.Printf("You connected to your %s database.\n", mongoDB)
}
