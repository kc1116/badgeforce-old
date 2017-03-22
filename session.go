package main

import (
	"errors"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//Login . . . check user password and get a session token
func NewSession(email string, password string) (Session, error) {
	var session Session
	user, err := checkPassword(email, password)
	if err != nil {
		err = errors.New("The email/password combination was not found.")
		return session, err
	}

	session, err = getSessionToken(user)
	if err != nil {
		err = errors.New("Session could not be established at this time try again later.")
		return session, err
	}

	return session, nil
}

func checkPassword(email string, password string) (BadgeforceUser, error) {
	var user BadgeforceUser
	user, err := GetUser(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

type Session struct {
	SessionToken string `json:"sessionToken"`
}

//BadgeForceJWTClaims . . .
type BadgeForceJWTClaims struct {
	User BadgeforceUser `json:"user"`
	jwt.StandardClaims
}

func getSessionToken(user BadgeforceUser) (Session, error) {
	var session Session
	mySigningKey := []byte(Config.App.TokenKey)

	claims := BadgeForceJWTClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return session, err
	}

	session.SessionToken = signedToken

	return session, nil
}

var verifySessionMW *jwtmiddleware.JWTMiddleware
