package model

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Injecting Data in html form
type Data struct {
	HomeLink      string
	CsrfToken     string
	LoginEndpoint string
}
