package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

// DataStore containing a pointer to a mgo session
type DataStore struct {
	Session *mgo.Session
}

//Store . . .
var Store DataStore

const (
	UserCollection = "BadgeForceUsers"
)

func init() {
	fmt.Print("Initializing MongoDB . . . . .10%")
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Config.Database.Addrs,
		Username: Config.Database.Username,
		Password: Config.Database.Password,
		Database: Config.Database.Database,
		Timeout:  time.Minute,
	})
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	Store.Session = session
	fmt.Print(". . . . .30%")
	userIndex := []mgo.Index{
		mgo.Index{
			Key:        []string{"system.uuid"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		},
		mgo.Index{
			Key:        []string{"email"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		},
		mgo.Index{
			Key:        []string{"firstname", "lastname"},
			Background: true,
			Sparse:     true,
		},
	}

	fmt.Print(". . . . .60%")
	ensureMgoIndexes(userIndex, UserCollection)
	fmt.Print(". . . . .100%\n")
}

func ensureMgoIndexes(indexes []mgo.Index, collection string) {
	tempSession := Store.Session.Copy()
	for _, i := range indexes {
		err := tempSession.DB(Config.Database.Database).C(collection).EnsureIndex(i)
		if err != nil {
			panic(err.Error())
		}
	}
}

//GetStore . . . returns a session store
func GetStore() *mgo.Session {
	tempStore := Store.Session.Copy()
	return tempStore
}
