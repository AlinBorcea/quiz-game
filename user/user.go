package user

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"time"
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

	file, err := os.OpenFile(filename, os.O_APPEND, fs.ModeAppend)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	err = writer.Write(userAsRecord(user))
	if err != nil {
		return err
	}

	writer.Flush()
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

func userAsRecord(user *User) []string {
	rec := make([]string, 3)
	format := "2006-01-02 15:04:05"

	rec[0] = user.Name
	rec[1] = time.Now().Format(format)
	rec[2] = strconv.Itoa(user.Score)
	return rec
}
