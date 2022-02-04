package quiz

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// A Quiz variable contains the records of a quiz.
type Quiz struct {
	records [][]string
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

// To be modified.
func AnswerIsCorrect(q *Question, answer int) bool {
	return q.correct == answer
}

// To be modified.
func (q *Quiz) QuestionAt(index int) (que *Question, err error) {
	if !(index >= 1 && index < len(q.records)) {
		return nil, fmt.Errorf("index %d is out of range 1..%d", index, len(q.records)-1)
	}

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
