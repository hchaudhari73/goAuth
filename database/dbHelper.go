package database

import (
	"errors"

	"github.com/hchaudhari73/goAuth/model"
)

// Queries user table and return the response matching the request
func CheckCredsWhileLogin(u *model.User) (*model.User, error) {
	responseUser := model.User{}
	if DB.Where(u).First(&responseUser).Error != nil {
		return nil, errors.New("error while getting user")
	}
	return &responseUser, nil
}
