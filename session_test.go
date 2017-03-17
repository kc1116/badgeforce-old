package main

import (
	"testing"
	"time"
)

var mockUser = BadgeforceUser{
	FirstName: string(getNewUserSalt()),
	LastName:  string(getNewUserSalt()),
	Email:     string(getNewUserSalt()),
	Password:  string(getNewUserSalt()),
	System: system{
		UUID:      "test",
		CreatedOn: time.Now(),
	},
}

/*func TestCheckPassword(t *testing.T) {

	actualUser, err := CreateUser(mockUser.FirstName, mockUser.LastName, mockUser.Email, "password")
	if err != nil {
		t.Error("Expected an user object got an error:", err.Error())
	}

	err = StoreUser(actualUser)
	if err != nil {
		t.Error("Attempted to store use got an error:", err.Error())
	}

	user, err := checkPassword(actualUser.System.UUID, "password")
	if err != nil {
		t.Error("TestCheckPassword error: ", err.Error())
	} else {
		t.Logf("TestCheckPassword: %v", user)
	}
}*/

func TestGetSession(t *testing.T) {
	session, err := getSessionToken(mockUser)
	if err != nil {
		t.Error("TestGetSession error: ", err.Error())
	} else {
		t.Logf("TestCheckPassword: %v", session.SessionToken)
	}
}
