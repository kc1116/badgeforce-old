package main

import (
	"math/rand"
	"time"

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
func CreateUser(fname string, lname string, email string, password string) (BadgeforceUser, userSalt, error) {
	uuid := GetUUID()
	salt := getNewUserSalt(uuid)
	hashedPword, err := hashPassword(password, salt)
	if err != nil {
		return BadgeforceUser{}, userSalt{}, err
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
	}, salt, nil
}

//StoreUser . . . create new user in Mongodb
func StoreUser(user BadgeforceUser, pwdSalt userSalt) error {
	store := GetStore()

	coll := store.DB(Config.Database.Database).C(UserCollection)
	if err := coll.Insert(user); err != nil {
		return err
	}
	coll = store.DB(Config.Database.Database).C(SaltCollection)
	if err := coll.Insert(pwdSalt); err != nil {
		return err
	}

	return nil
}

//GetUUID . . .returns a uuid v4 string
func GetUUID() string {
	return uuid.NewV4().String()
}

func hashPassword(password string, salt userSalt) (string, error) {
	hashedPwrd, err := bcrypt.GenerateFromPassword([]byte(salt.Salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwrd), nil
}

func getNewUserSalt(uuid string) userSalt {
	b := make([]byte, 64)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return userSalt{
		User: uuid,
		Salt: string(b),
	}
}
