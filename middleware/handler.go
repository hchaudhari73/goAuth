package middleware

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/hchaudhari73/goAuth/config"
	"github.com/hchaudhari73/goAuth/model"
)

var (
	home     = "templates/home.html"
	login    = "templates/login.html"
	userhome = "templates/userhome.html"
)

// Landing page for your api
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: Home")

	// Setting CSRF header
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	// Parsing html file
	parsedTemplate, err := template.ParseFiles(home)
	if err != nil {
		fmt.Println(err)
	}

	// Link to Login page
	loginEP, err := getLoginEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	data := model.Data{
		HomeLink: *loginEP,
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

/*
	Login
	Get: This is a get function to load html and inject csrf

*/
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: Login")

	if r.Method == "POST" {
		// For POST method
		// Getting data from request
		var user model.User
		json.NewDecoder(r.Body).Decode(&user)

		// Verify user

		if "" == user.Email && "" == user.Password {

			/* session */
			sessionToken, err := config.GetSessionToken()
			if err != nil {
				fmt.Println(err)
			}
			store := sessions.NewCookieStore([]byte(*sessionToken))
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
				Name:    "email",
				Value:   fmt.Sprint(session.Values["email"]),
				Path:    "/",
				Expires: time.Now().Add(time.Hour),

				HttpOnly: true,
				Secure:   true,
			}
			http.SetCookie(w, &cookie)

		} else {
			fmt.Fprintln(w, http.StatusBadRequest, map[string]string{"msg": "Invalid Payload"})
		}
	}

	// For GET method
	// Setting csrf token
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	// Parsing html
	parsedTemplate, err := template.ParseFiles(login)
	if err != nil {
		fmt.Println(err)
	}

	// Get links
	homeEP, err := getHomeEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	loginEP, err := getLoginEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	data := model.Data{
		HomeLink:      *homeEP,
		CsrfToken:     csrf.Token(r),
		LoginEndpoint: *loginEP,
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

/*
	UserHome
	To display user home page, after successful signin
*/
func UserHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: UserHome")

	/*
		First need to check weather the user's email
		from cookies is present.
	*/
	// Getting cookies
	cookies := r.Cookies()

	// Getting our user data
	if len(cookies) < 1 {
		// Redirecting to home
		homeEP, err := getHomeEndpoint()
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, *homeEP, http.StatusPermanentRedirect)
	}

	/* Parsing html */
	parsedTemplate, err := template.ParseFiles(userhome)
	if err != nil {
		fmt.Println(err)
	}

	homeEP, err := getHomeEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	data := model.Data{
		HomeLink: *homeEP,
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
