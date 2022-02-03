package quiz

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	records [][]string
}

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

func AnswerIsCorrect(q *Question, answer int) bool {
	return q.correct == answer
}

func (q *Quiz) GetQuestion(index int) (que *Question, err error) {
	if !(index >= 1 && index < len(q.records)) {
		return nil, fmt.Errorf("index %d is out of range 1..%d", index, len(q.records)-1)
	}

	return fromRecord(q.records[index])
}

func (q *Quiz) GetRandom() (que *Question, err error) {
	rand.Seed(time.Now().Unix())
	index := rand.Int() % len(q.records)
	if index == 0 {
		index = 1
	}
	return q.GetQuestion(index)
}

func readRecords(filename string) (records [][]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	records, err = csv.NewReader(file).ReadAll()
	file.Close()

	return records, err
}
