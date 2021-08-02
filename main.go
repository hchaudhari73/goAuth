package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/hchaudhari73/goAuth/router"
)

const (
	base      = "localhost:8080"
	csrfToken = "t3rc3s-p0t"
)

func main() {
	router := router.Router()

	// csrf
	csrfRouter := csrf.Protect(
		[]byte(csrfToken),
		csrf.Path("/"),
		csrf.Secure(false),
	)(router)

	// init server
	fmt.Printf("goAuth running on %s\n", base)
	http.ListenAndServe(base, csrfRouter)
}
