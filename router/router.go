package router

import (
	"github.com/gorilla/mux"
	"github.com/hchaudhari73/goAuth/middleware"
)

// router
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Home).Methods("GET")
	r.HandleFunc("/login", middleware.Login).Methods("POST")
	return r
}
