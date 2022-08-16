package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// server main method

var router = mux.NewRouter()

var redirectTarget string

func main() {
	
	router.HandleFunc("/", index)
	router.HandleFunc("/signinpage", signinPage)
	router.HandleFunc("/signin", signinHandler).Methods("POST")
	router.HandleFunc("/loginpage", loginPage)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/{Uname}", internalPageHandler)
	router.HandleFunc("/{Uname}/updatepass", updatepassHandler).Methods("POST")
	router.HandleFunc("/{Uname}/dbpage", dbpage) //create assignment page open
	router.HandleFunc("/{Uname}/addnew", addNewAssign).Methods("POST") // create assignment	
	router.HandleFunc("/{Uname}/assignment", assignmentPage) // observe assignments
	router.HandleFunc("/{Uname}/assuppage", updateAssPageHandler) //open update page
	router.HandleFunc("/{Uname}/assupdate", updateAssHandler).Methods("POST") // update status
	router.HandleFunc("/{Uname}/logout", logoutHandler).Methods("POST")
	router.HandleFunc("/{Uname}/delete", deleteHandler).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
