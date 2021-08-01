package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/hchaudhari73/goAuth/router"
)

const base = "localhost:8080"
const csrfToken = "t3rc3s-p0t"

func main() {
	router := router.Router()
	fmt.Printf("goAuth running on %s\n", base)
	csrfRouter := csrf.Protect([]byte(csrfToken))(router)
	http.ListenAndServe(base, csrfRouter)
}
