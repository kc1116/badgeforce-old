package main

import (
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func checkPassword(uuid string, password string) (BadgeforceUser, error) {
	var user BadgeforceUser
	user, err := GetUser(uuid)
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
	mySigningKey := []byte("AllYourBase")

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

var verifySessionMW jwtmiddleware.JWTMiddleware
