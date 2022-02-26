package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AlinBorcea/quiz-game/quiz"
)

const (
	queFilename   = "questions.csv"
	scoreFilename = "scores.csv"
)

func main() {
	q, err := quiz.New(queFilename)
	if err != nil {
		log.Fatalln(err)
	}

	runTestRandom(q)

	reader := bufio.NewReader(os.Stdin)
	nameSlice, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}

	name := string(nameSlice)
	name = strings.TrimSuffix(name, "\n")
	name = strings.TrimSuffix(name, "\r")

	postResult(name, q.Result())
}

func runTestRandom(q *quiz.Quiz) {
	var ans int

	reader := bufio.NewReader(os.Stdin)
	que, err := q.RandomQuestion()
	for err == nil {
		printQuestion(&que)
		ans, _ = readInput(reader)
		q.Answer(ans)

		que, err = q.RandomQuestion()
	}
}

func printQuestion(q *quiz.Question) {
	fmt.Println(q.Que)
	for i := 0; i < len(q.Answers); i++ {
		fmt.Print(q.Answers[i])
		if i%2 == 0 {
			fmt.Print("\t")
		} else {
			fmt.Println()
		}
	}
	if len(q.Answers)%2 == 0 {
		fmt.Print("\n")
	} else {
		fmt.Printf("\n\n")
	}
}

func readInput(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	ans, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return ans, nil
}

func postResult(name string, result int) {
	fmt.Printf("%v -> %d\n", name, result)
}
