package db

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// DB session
var DB *mgo.Database

// Games collection
var Games *mgo.Collection

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

	fmt.Printf("You connected to your %s database.\n", mongoDB)
}
