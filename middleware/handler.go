package middleware

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/csrf"
	"github.com/hchaudhari73/goAuth/config"
	"github.com/hchaudhari73/goAuth/database"
	"github.com/hchaudhari73/goAuth/model"
)

var (
	home     = "templates/home.html"
	login    = "templates/login.html"
	userhome = "templates/userhome.html"
	notFound = "templates/four04.html"
)

// Landing page for your api
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: Home")

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

	if isEmailPresent(r) {
		http.Redirect(w, r, "/userhome", http.StatusPermanentRedirect)
	}

	if r.Method == "POST" {
		// For POST method
		// Getting data from request
		var user model.User
		json.NewDecoder(r.Body).Decode(&user)

		// Verify user
		_, err := database.CheckCredsWhileLogin(&user)
		if err != nil {
			parsedTemp, err := template.ParseFiles(notFound)
			if err != nil {
				fmt.Println("Error while parsing 404 error", err)
			}
			parsedTemp.Execute(w, nil)
			return
		}

		jwtExpiry := time.Now().Add(8 * time.Hour)
		claim := jwt.StandardClaims{
			ExpiresAt: jwtExpiry.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		key, err := config.GetJWTkey()
		if err != nil {
			fmt.Println("Error while getting JWT key:", err)
			return
		}
		tokenString, err := token.SignedString(key)
		if err != nil {
			fmt.Println("Error while creating string token:", err)
			fmt.Fprintf(w, "internal server error")
			return
		}

		// Setting cookies
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: jwtExpiry,
		})
	}

	// For GET method
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

	/*
		First need to check weather the user's email
		from cookies is present.
	*/
	// Checking if session is present in cookies
	if !isJWTTokenPresent(r) {
		// Redirecting to home
		homeEP, err := getHomeEndpoint()
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, *homeEP, http.StatusPermanentRedirect)
	}

	fmt.Println("Endpoint hit: UserHome")

	/* Parsing html */
	parsedTemplate, err := template.ParseFiles(userhome)
	if err != nil {
		fmt.Println(err)
	}

	homeEP, err := getHomeEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	logoutEP, err := getLogoutEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	data := model.Data{
		HomeLink:       *homeEP,
		LogoutEndpoint: *logoutEP,
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

// Logout
func Logout(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Logout")

	// Removing `token` cookie from the browser
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})

	homeEP, err := getHomeEndpoint()
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, *homeEP, http.StatusPermanentRedirect)
}
