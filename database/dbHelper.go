package database

import (
	"github.com/hchaudhari73/goAuth/model"
)

// Queries user table and return the response matching the request
func CheckCredsWhileLogin(u *model.User) *model.User {
	responseUser := model.User{}
	DB.Where(u).First(&responseUser)
	return &responseUser
}
