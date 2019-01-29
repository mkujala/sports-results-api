package db

import (
	"fmt"
	"sports-results/config"
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
	config := config.Values()
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.DB.URI},
		Timeout:  60 * time.Second,
		Database: config.DB.Name,
		Username: config.DB.User,
		Password: config.DB.Pass,
	}

	s, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB(config.DB.Name)
	Games = DB.C("games")
	Standings = DB.C("standings")

	fmt.Printf("You connected to your %s database.\n", config.DB.Name)
}
