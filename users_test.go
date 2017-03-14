package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	pword, err := hashPassword("password")
	if err != nil {
		t.Error("Expected an user object got an error:", err.Error())
	}

	expectedUser := BadgeforceUser{
		FirstName: "Khalil",
		LastName:  "Claybon",
		Email:     "kc@gmail.com",
		Password:  pword,
		System: system{
			UUID:      "test",
			CreatedOn: time.Now(),
		},
	}

	actualUser, err := CreateUser(expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, "password")
	if err != nil {
		t.Error("Expected an user object got an error:", err.Error())
	}
	fmt.Printf("%s\n", expectedUser.Password)
	fmt.Printf("%s\n", actualUser.Password)
}
