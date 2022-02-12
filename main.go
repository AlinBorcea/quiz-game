package main

import (
	"fmt"
	"log"

	"github.com/AlinBorcea/quiz-game/quiz"
)

const filename = "questions.csv"

func main() {
	q, err := quiz.New(filename)
	if err != nil {
		log.Fatalln(err)
	}

	x := 0
	goodAnsers := 0
	for i := 0; i < 5; i++ {
		que, err := q.RandomQuestion()
		if err != nil {
			log.Fatalln(err)
		}
		printQuestion(que)
		x = 4
		if q.Answer(x) {
			goodAnsers++
		}
	}

	fmt.Printf("%d good answers", goodAnsers)

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
