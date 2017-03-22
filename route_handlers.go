package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type FormValidation interface {
	Validate() error
}

type result struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//RegisterRes . . . http response struct for the register route
type RegisterRes struct {
	Session Session `json:"session"`
	Result  result  `json:"result"`
}

type LoginRes struct {
	Session Session `json:"session"`
	Result  result  `json:"result"`
}

//Register . . . register new user
func Register(w http.ResponseWriter, r *http.Request) {

	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	session, err := NewUser(fname, lname, email, password)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Registration could not be completed at this time.", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(RegisterRes{
		Session: session,
		Result: result{
			Error:   "",
			Message: "Registration successful",
			Status:  200,
		},
	}); err != nil {
		log.Println("JSON encoder error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Login . . .
func Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	session, err := NewSession(email, password)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if err = json.NewEncoder(w).Encode(LoginRes{Session: session}); err != nil {
		log.Println("JSON encoder error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
