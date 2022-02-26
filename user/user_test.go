package user

import "testing"

const scoreFilename = "scores.csv"

func TestPostUserScore_EmptyUser(t *testing.T) {
	user := User{Name: "", Score: 0}
	if err := PostUserScore(&user, scoreFilename); err == nil {
		t.Fatal("empty user is invalid")
	}
}

func TestPostUserScore_NegativeScore(t *testing.T) {
	user := User{Name: "alin", Score: -1}
	if err := PostUserScore(&user, scoreFilename); err == nil {
		t.Fatal("user score cannot be a negative integer")
	}
}

func TestPostUserScore_EmptyFilename(t *testing.T) {
	user := User{Name: "das", Score: 0}
	if err := PostUserScore(&user, ""); err == nil {
		t.Fatal("empty filename is invalid")
	}
}

func TestPostUserScore_NotACsvFile(t *testing.T) {
	user := User{Name: "asd", Score: 0}
	if err := PostUserScore(&user, "some.txt"); err == nil {
		t.Fatal("file is not a .csv file")
	}
}

func Test_userIsValid(t *testing.T) {

}
