package router

import (
	"github.com/gorilla/mux"
	"github.com/hchaudhari73/goAuth/middleware"
)

// router
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Home).Methods("GET")
	r.HandleFunc("/home", middleware.Base).Methods("GET")
	r.HandleFunc("/login", middleware.Login).Methods("GET", "POST") // To load the html and inject csrf token
	r.HandleFunc("/userhome", middleware.UserHome).Methods("GET")   // To load html for user after successful login.
	r.HandleFunc("/logout", middleware.Logout).Methods("GET")       // To logout
	return r
}
