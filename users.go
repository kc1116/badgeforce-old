package main

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

const (
	collection = "BadgeForceUsers"
)

//BadgeforceUser . . .
type BadgeforceUser struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	System    system
}

type system struct {
	UUID      string
	CreatedOn time.Time
}

//CreateUser . . . return proper user struct
func CreateUser(fname string, lname string, email string, password string) BadgeforceUser {

}

//StoreUser . . . create new user in Mongodb
func StoreUser(user BadgeforceUser) error {
	store := GetStore()
	coll := store.DB(Config.Database.Database).C(collection)
	if err := coll.Insert(user); err != nil {
		return err
	}
	return nil
}

//Initialize indexes in mongodb for BadgeForceUsers
func init() {
	store := GetStore()
	defer store.Close()
	index := []mgo.Index{
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
	for _, i := range index {
		err := store.DB(Config.Database.Database).C(collection).EnsureIndex(i)
		if err != nil {
			panic(err)
		}
	}
}
