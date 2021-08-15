package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/hchaudhari73/goAuth/config"
	"github.com/hchaudhari73/goAuth/router"
)

func main() {
	router := router.Router()

	// csrf
	csrfToken, err := config.GetCsrfToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	csrfRouter := csrf.Protect(
		[]byte(*csrfToken),
		csrf.Path("/"),
		csrf.MaxAge(3600),
		csrf.Secure(false),
	)(router)

	// Getting server parameters
	base, err := config.GetBaseUrl()
	if err != nil {
		fmt.Println(err)
		return
	}

	port, err := config.GetPort()
	if err != nil {
		fmt.Println(err)
		return
	}

	runOn := fmt.Sprintf("%s:%s", *base, *port)

	// init server
	fmt.Printf("goAuth running on %s\n", runOn)
	http.ListenAndServe(runOn, csrfRouter)
}
