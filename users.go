package main

import (
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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

type userSalt struct {
	User string `bson:"user"`
	Salt string `bson:"salt"`
}

//CreateUser . . . return proper user struct
func CreateUser(fname string, lname string, email string, password string) (BadgeforceUser, error) {
	uuid := GetUUID()
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
			UUID:      uuid,
			CreatedOn: time.Now(),
		},
	}, nil
}

//StoreUser . . . create new user in Mongodb
func StoreUser(user BadgeforceUser) error {
	store := GetStore()
	defer store.Close()
	coll := store.DB(Config.Database.Database).C(UserCollection)
	if err := coll.Insert(user); err != nil {
		return err
	}

	return nil
}

//GetUser . . . gets a user from the database based on user uuid
func GetUser(uuid string) (BadgeforceUser, error) {
	store := GetStore()
	defer store.Close()
	coll := store.DB(Config.Database.Database).C(UserCollection)

	user := BadgeforceUser{}
	err := coll.Find(bson.M{"system.uuid": uuid}).One(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

//GetUUID . . .returns a uuid v4 string
func GetUUID() string {
	return uuid.NewV4().String()
}

func hashPassword(password string) (string, error) {
	hashedPwrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwrd), nil
}

func getNewUserSalt() []byte {
	b := make([]byte, 64)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
