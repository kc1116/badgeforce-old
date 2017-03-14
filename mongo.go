package main

import "gopkg.in/mgo.v2"
import "time"

// DataStore containing a pointer to a mgo session
type DataStore struct {
	Session *mgo.Session
}

//Store . . .
var Store DataStore

func init() {
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

	userIndex := []mgo.Index{
		mgo.Index{
			Key:        []string{"system.uuid", "email"},
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
	for _, i := range userIndex {
		err := session.DB(Config.Database.Database).C(collection).EnsureIndex(i)
		if err != nil {
			panic(err)
		}
	}
}

//GetStore . . . returns a session store
func GetStore() *mgo.Session {
	tempStore := Store.Session.Copy()
	return tempStore
}
