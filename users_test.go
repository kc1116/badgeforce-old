package main

import (
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	expectedUser := BadgeforceUser{
		FirstName: "Khalil",
		LastName:  "Claybon",
		Email:     "kc@gmail.com",
		Password:  "password",
		System: system{
			UUID:      "test",
			CreatedOn: time.Now(),
		},
	}

	actualUser, err := CreateUser(expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, "password")
	if err != nil {
		t.Error("Expected an user object got an error:", err.Error())
	} else {
		t.Logf("TestCreateUser:%v", actualUser)
	}
}

/*func TestStoreUser(t *testing.T) {
	actualUser, err := CreateUser("John", "Doe", "jd@gmail.com", "password")
	if err != nil {
		t.Error("Expected an user object got an error:", err.Error())
	}

	err = StoreUser(actualUser)
	if err != nil {
		t.Error("Attempted to store use got an error:", err.Error())
	} else {
		t.Logf("TestStoreUser:%v", actualUser)
	}

}*/
