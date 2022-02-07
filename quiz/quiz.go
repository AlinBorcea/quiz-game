package quiz

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// A Quiz variable contains the records of a quiz.
type Quiz struct {
	records       [][]string
	currentRecord int
}

// A Question is an easier to use representation of a quiz record.
type Question struct {
	Que     string
	Answers []string
}

// New tries to create a Quiz variable and returns it if no errors
// were encountered while reading the records from the specified file.
// filename must include the .csv extension.
func New(filename string) (*Quiz, error) {
	if !strings.Contains(filename, ".csv") {
		return nil, fmt.Errorf("file `%s` is not a csv file", filename)
	}

	records, err := readRecords(filename)
	if err != nil {
		return nil, err
	}

	if !(len(records) > 1) {
		return nil, fmt.Errorf("there are no records in file `%s`", filename)
	}

	return &Quiz{records: records}, nil
}

func (q *Quiz) QuestionAt(index int) (que *Question, err error) {
	if !(index >= 1 && index < len(q.records)) {
		return nil, fmt.Errorf("index %d is out of range 1..%d", index, len(q.records)-1)
	}

	q.currentRecord = index
	return questionFromRecord(q.records[index])
}

// RandomQuestion uses rand to generate a random index and returns
// the result of a call to QuestionAt.
func (q *Quiz) RandomQuestion() (que *Question, err error) {
	rand.Seed(time.Now().Unix())
	index := rand.Int() % len(q.records)
	if index == 0 {
		index = 1
	}

	return q.QuestionAt(index)
}

func (q *Quiz) AnswerIsCorrect(answer int) bool {
	correct, err := strconv.Atoi(q.records[q.currentRecord][2])
	if err != nil || answer != correct {
		return false
	}
	return true
}

// readRecords takes a filename and tries to read all the records of the file.
// The file should be a csv file otherwise the operation might fail.
func readRecords(filename string) (records [][]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	records, err = csv.NewReader(file).ReadAll()
	file.Close()

	return records, err
}

// questionFromRecord takes a quiz record, checks if it is valid and returns an easier to use variable.
func questionFromRecord(record []string) (*Question, error) {
	if err := recordIsValid(record); err != nil {
		return nil, err
	}

	return &Question{record[0], strings.Split(record[1], ",")}, nil
}

func recordIsValid(record []string) error {
	que := record[0]
	if len(que) <= 0 {
		return fmt.Errorf("record question is empty. %d", len(que))
	}

	if len(record[1]) <= 0 {
		return errors.New("record has no answers")
	}

	answers := strings.Split(record[1], ",")
	if hasEmptyAnswer(answers) {
		return errors.New("an answer has invalid length")
	}

	if _, err := strconv.Atoi(record[2]); err != nil {
		return err
	}

	return nil
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
