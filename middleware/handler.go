package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/hchaudhari73/goAuth/data"
)

// This is used to sign the token
var store = sessions.NewCookieStore([]byte("t3rc3s-p0t"))

// This is the landing page for your api
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: Home")

	/* Setting CSRF header  */
	w.Header().Set("X-CSRF-TOKEN", csrf.Token(r))

	fmt.Fprint(w, "Welcome to goAuth!")
}

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: Login")

	var user data.User
	json.NewDecoder(r.Body).Decode(&user)

	/* session */
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}

	// saving session variables
	session.Values["email"] = user.Email
	session.Save(r, w)

	/*
		Server side cookies
		Also will be able to test sessions
	*/
	cookie := http.Cookie{
		Name:  "email",
		Value: fmt.Sprint(session.Values["email"]),
		Path:  "/",

		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	fmt.Println("Cookies set")

	w.Header().Set("X-CSRF-TOKEN", csrf.Token(r))

	fmt.Fprintln(w, "Logged in successfully")
}
