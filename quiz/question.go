package quiz

import (
	"fmt"
	"strconv"
	"strings"
)

// A Question is an easier to use representation of a quiz record.
// Only the Que and Answers fields are exported while the correct
// answer is unexported to prevent abuse.
type Question struct {
	Que     string
	Answers []string
	correct int
}

// questionFromRecord takes a quiz record, checks if it is valid and returns an easier to use variable.
func questionFromRecord(record []string) (*Question, error) {
	que := record[0]
	if len(que) <= 0 {
		return nil, fmt.Errorf("field Que of Question is empty. %d", len(que))
	}

	if len(record[1]) <= 0 {
		return nil, fmt.Errorf("there are no answers to this question")
	}

	answers := strings.Split(record[1], ",")
	if hasEmptyAnswer(answers) {
		return nil, fmt.Errorf("an answer has invalid length")
	}

	correct, err := strconv.Atoi(record[2])
	if err != nil {
		return nil, err
	}

	return &Question{Que: que, Answers: answers, correct: correct}, nil
}

// hasEmptyAnswer reports if there is an empty string in a []string.
func hasEmptyAnswer(answers []string) bool {
	for i := 0; i < len(answers); i++ {
		if len(answers[i]) <= 0 {
			return true
		}
	}
	return false
}
