package db

import (
	"fmt"
	"time"

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
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoURI},
		Timeout:  60 * time.Second,
		Database: mongoDB,
		Username: mongoUser,
		Password: mongoPass,
	}

	s, err := mgo.DialWithInfo(mongoDBDialInfo)
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
