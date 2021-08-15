package data

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Injecting Data in html form
type Data struct {
	Link          string
	CsrfToken     string
	LoginEndpoint string
}

// Sample Data
var user = User{
	Name:     "Dummy User",
	Email:    "dummy@email.com",
	Password: "dummyPassword",
}

// user info getter
func GetData() *User {
	return &user
}
