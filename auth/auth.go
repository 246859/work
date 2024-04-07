package auth

import (
	"errors"
	"github.com/246859/work/user"
)

// Verify verifies user credentials is ok
func Verify(user user.User) (bool, error) {
	password, err := query(user.Name)
	if err != nil {
		return false, err
	}
	if password != user.Password {
		return false, errors.New("authentication failed")
	}
	return true, nil
}

func query(username string) (string, error) {
	if username == "jack" {
		return "jack123456", nil
	}
	return "", errors.New("user not found")
}
