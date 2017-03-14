package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
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
func CreateUser(fname string, lname string, email string, password string) (BadgeforceUser, error) {
	hashedPword, err := hashPassword(password)
	if err != nil {
		return BadgeforceUser{}, err
	}
	return BadgeforceUser{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  hashedPword,
		System: system{
			UUID:      GetUUID(),
			CreatedOn: time.Now(),
		},
	}, nil
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

//GetUUID . . .returns a uuid v4 string
func GetUUID() string {
	return uuid.NewV4().String()
}

func hashPassword(password string) (string, error) {
	hashedPwrd, err := bcrypt.GenerateFromPassword([]byte(password+Config.App.Salt), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword(hashedPwrd, []byte(password+Config.App.Salt))
	if err != nil {
		panic(err)
	}
	return string(hashedPwrd), nil
}
