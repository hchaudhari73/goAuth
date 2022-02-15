package model

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Checks if the Name of the
// user is present or not
func (u *User) IsValid() bool {
	// valid if the name is not empty
	// else false
	return u.Name != ""
}

// Injecting Data in html form
type Data struct {
	HomeLink      string
	CsrfToken     string
	LoginEndpoint string
}
