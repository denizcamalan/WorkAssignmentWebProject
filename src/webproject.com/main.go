package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// server main method

var router = mux.NewRouter()

var redirectTarget string

func main() {
	
	router.HandleFunc("/denizcamalan.github.io/", index)
	router.HandleFunc("/denizcamalan.github.io/signinpage", signinPage)
	router.HandleFunc("/denizcamalan.github.io/signin", signinHandler).Methods("POST")
	router.HandleFunc("/denizcamalan.github.io/loginpage", loginPage)
	router.HandleFunc("/denizcamalan.github.io/login", loginHandler).Methods("POST")
	router.HandleFunc("/denizcamalan.github.io/{Uname}", internalPageHandler)
	router.HandleFunc("/denizcamalan.github.io/{Uname}/updatepass", updatepassHandler).Methods("POST")
	router.HandleFunc("/denizcamalan.github.io/{Uname}/dbpage", dbpage) //create assignment page open
	router.HandleFunc("/denizcamalan.github.io/{Uname}/addnew", addNewAssign).Methods("POST") // create assignment	
	router.HandleFunc("/denizcamalan.github.io/{Uname}/assignment", assignmentPage) // observe assignments
	router.HandleFunc("/denizcamalan.github.io/{Uname}/assuppage", updateAssPageHandler) //open update page
	router.HandleFunc("/denizcamalan.github.io/{Uname}/assupdate", updateAssHandler).Methods("POST") // update status
	router.HandleFunc("/denizcamalan.github.io/{Uname}/logout", logoutHandler).Methods("POST")
	router.HandleFunc("/denizcamalan.github.io/{Uname}/delete", deleteHandler).Methods("POST")
	http.Handle("/denizcamalan.github.io/", router)
	http.ListenAndServe(":8080", nil)
}
