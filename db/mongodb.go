package db

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var (
	// DB session
	DB *mgo.Database

	// Games Collection
	Games *mgo.Collection

	// Standings Collection
	Standings *mgo.Collection
)

func init() {
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
