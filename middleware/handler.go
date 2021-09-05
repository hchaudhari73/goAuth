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
	"github.com/hchaudhari73/goAuth/data"
)

// This is used to sign the token
var (
	home     = "templates/home.html"
	login    = "templates/login.html"
	userhome = "templates/userhome.html"
)

// This is the landing page for your api
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: Home")

	/* Setting CSRF header  */
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	// Parsing html file
	parsedTemplate, err := template.ParseFiles(home)
	if err != nil {
		fmt.Println(err)
	}

	// Link to Login page
	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		fmt.Println(err)
	}

	port, err := config.GetPort()
	if err != nil {
		fmt.Println(err)
	}

	data := data.Data{
		Link: fmt.Sprintf("%s:%s/login", *baseHttp, *port),
	}
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

/*
	Login
	This is a get function to load html and inject csrf
*/
func LoginGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: LoginGet")

	// Setting csrf token
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	/* Parsing html */
	parsedTemplate, err := template.ParseFiles(login)
	if err != nil {
		fmt.Println(err)
	}

	// Parsing link for homepage
	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		fmt.Println(err)
	}

	port, err := config.GetPort()
	if err != nil {
		fmt.Println(err)
	}

	data := data.Data{
		Link:          fmt.Sprintf("%s:%s", *baseHttp, *port),
		CsrfToken:     csrf.Token(r),
		LoginEndpoint: fmt.Sprintf("%s:%s/login/post", *baseHttp, *port),
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

/*
	Login Post
	post api to send user data to login
*/
func LoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: LoginPost")

	// Getting data from request
	var user data.User
	json.NewDecoder(r.Body).Decode(&user)

	// Verify user
	inDataUser := data.GetData()
	if inDataUser.Email == user.Email && inDataUser.Password == user.Password {

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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, map[string]string{"response": "Invalid playload"})
	}

}

/*
	UserGet
	To display user home page, after successful signin
*/
func UserGet(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: UserGet")

	/* Parsing html */
	parsedTemplate, err := template.ParseFiles(userhome)
	if err != nil {
		fmt.Println(err)
	}

	// Parsing link for homepage
	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		fmt.Println(err)
	}

	port, err := config.GetPort()
	if err != nil {
		fmt.Println(err)
	}

	data := data.Data{
		Link: fmt.Sprintf("%s:%s", *baseHttp, *port),
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
