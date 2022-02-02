package quiz

import (
	"fmt"
	"strconv"
	"strings"
)

type Question struct {
	Que     string
	Answers []string
	Correct int
}

func fromRecord(record []string) (*Question, error) {
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

	return &Question{Que: que, Answers: answers, Correct: correct}, nil
}

func hasEmptyAnswer(answers []string) bool {
	for i := 0; i < len(answers); i++ {
		if len(answers[i]) <= 0 {
			return true
		}
	}
	return false
}
