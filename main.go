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

const filename = "questions.csv"

func main() {
	q, err := quiz.New(filename)
	if err != nil {
		log.Fatalln(err)
	}

	runTest(q)
}

func runTest(q *quiz.Quiz) {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < q.Len(); i++ {
		que, err := q.QuestionAt(i)
		if err != nil {
			break
		}

		printQuestion(que)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")
		ans, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		q.Answer(ans)

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
