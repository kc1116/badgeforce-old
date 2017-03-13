package main

import "gopkg.in/mgo.v2"

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
	})
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	Store.Session = session

}

//GetStore . . . returns a session store
func GetStore() *mgo.Session {
	tempStore := Store.Session.Copy()
	return tempStore
}
