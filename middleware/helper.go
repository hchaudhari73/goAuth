package middleware

import (
	"fmt"
	"net/http"

	"github.com/hchaudhari73/goAuth/config"
)

func getHomeEndpoint() (*string, error) {

	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		return nil, err
	}

	port, err := config.GetPort()
	if err != nil {
		return nil, err
	}

	// login endpoint
	endpoint := fmt.Sprintf("%s:%s", *baseHttp, *port)
	return &endpoint, nil
}

func getLoginEndpoint() (*string, error) {

	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		return nil, err
	}

	port, err := config.GetPort()
	if err != nil {
		return nil, err
	}

	// login endpoint
	endpoint := fmt.Sprintf("%s:%s/login", *baseHttp, *port)
	return &endpoint, nil
}

// Check if email is stored in the cookies
func isEmailPresent(r *http.Request) bool {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "email" && cookie.Value != "" {
			return true
		}
	}
	return false
}

// Check for `_gorilla_csrf` token in cookies
func isCSRFPresent(r *http.Request) bool {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "_gorilla_csrf" && cookie.Value != "" {
			return true
		}
	}
	return false
}
