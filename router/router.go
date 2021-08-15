package router

import (
	"github.com/gorilla/mux"
	"github.com/hchaudhari73/goAuth/middleware"
)

// router
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Home).Methods("GET")
	r.HandleFunc("/login", middleware.LoginGet).Methods("GET") // To load the html and inject csrf token
	r.HandleFunc("/login/post", middleware.LoginPost).Methods("POST")

	r.HandleFunc("/userhome", middleware.UserGet).Methods("GET") // To load html for user after successful login.
	return r
}
