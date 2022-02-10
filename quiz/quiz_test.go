package quiz

import "testing"

func TestHasEmptyAnswer_Empty(t *testing.T) {
	var answers []string

	if !hasEmptyAnswer(answers) {
		t.FailNow()
	}
}

func TestHasEmptyAnswer_1Empty(t *testing.T) {
	answers := []string{"1", "", "dsa"}
	if !hasEmptyAnswer(answers) {
		t.FailNow()
	}
}

func TestHasEmptyAnswer_AllEmpty(t *testing.T) {
	answers := []string{"", "", ""}
	if !hasEmptyAnswer(answers) {
		t.FailNow()
	}
}
