package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

func (q *Quiz) GetQuestion(index int) (question string, err error) {
	if !(index >= 1 && index < len(q.records)-1) {
		return "", fmt.Errorf("index %d is out of range 1..%d", index, len(q.records)-1)
	}

	return q.records[index][0], nil
}

func readRecords(filename string) (records [][]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(file).ReadAll()
}
