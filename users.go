package main

import (
	"log"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	"errors"

	"github.com/satori/go.uuid"
	"gopkg.in/asaskevich/govalidator.v4"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	usrPwrdRegex = "?!^[0-9]*$)(?!^[a-zA-Z]*$)^([a-zA-Z0-9]{8,10}"
)

//BadgeforceUser . . .
type BadgeforceUser struct {
	FirstName string `json:"firstname" bson:"firstname" valid:"alpha,required"`
	LastName  string `json:"lastname" bson:"lastname" valid:"alpha,required"`
	Email     string `json:"email" bson:"email" valid:"email,required"`
	Password  string `json:"password" bson:"password" valid:"alphanum, required"`
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

//NewUser . . . create user, store the user created, and get session token
func NewUser(fname string, lname string, email string, password string) (Session, error) {
	var session Session

	user, err := createUser(fname, lname, email, password)
	if err != nil {
		err = errors.New("Registration could not be completed at this time.")
		return session, err
	}

	err = user.save()
	if err != nil {
		log.Println(err.Error())
		err = errors.New("Registration could not be completed at this time.")
		return session, err
	}

	session, err = getSessionToken(user)
	if err != nil {
		err = errors.New("Session could not be established at this time try again later.")
		return session, err
	}

	return session, nil
}

//createUser . . . return proper user struct
func createUser(fname string, lname string, email string, password string) (BadgeforceUser, error) {
	uuid := GetUUID()
	user := BadgeforceUser{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  password,
		System: system{
			UUID:      uuid,
			CreatedOn: time.Now(),
		},
	}
	err := user.Validate()
	if err != nil {
		log.Println("HERE")
		err = errors.New("Registration could not be completed at this time.")
		return user, err
	}
	hashedPword, err := hashPassword(password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPword
	return user, nil
}

//Validate . . . user implementation of form validation
func (user *BadgeforceUser) Validate() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	return nil
}

//storeUser . . . create new user in Mongodb
func (user *BadgeforceUser) save() error {
	store := GetStore()
	defer store.Close()
	coll := store.DB(Config.Database.Database).C(UserCollection)
	if err := coll.Insert(user); err != nil {
		return err
	}

	return nil
}

//GetUser . . . gets a user from the database based on user uuid
func GetUser(email string) (BadgeforceUser, error) {
	store := GetStore()
	defer store.Close()
	coll := store.DB(Config.Database.Database).C(UserCollection)

	var user BadgeforceUser
	err := coll.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

//GetUser . . . gets a user from the database based on user uuid
func getUserByUUID(uuid string) (BadgeforceUser, error) {
	store := GetStore()
	defer store.Close()
	coll := store.DB(Config.Database.Database).C(UserCollection)

	var user BadgeforceUser
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

func init() {
	/*govalidator.TagMap["pword"] = govalidator.Validator(func(password string) bool {
		return regexp.MustCompile(usrPwrdRegex).MatchString(password)
	})*/
}
