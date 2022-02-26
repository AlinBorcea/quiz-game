package user

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Score int
}

func PostUserScore(user *User, filename string) error {
	if !userIsValid(user) {
		return fmt.Errorf("user %v is invalid", user)
	}

	if !strings.Contains(filename, ".csv") {
		return fmt.Errorf("file %s is not a csv file", filename)
	}

	return nil
}

func userIsValid(user *User) bool {
	if len(user.Name) <= 0 {
		return false
	}

	if user.Score < 0 {
		return false
	}

	return true
}
