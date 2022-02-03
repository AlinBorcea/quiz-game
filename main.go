package main

import (
	"fmt"
	"log"

	"github.com/AlinBorcea/quiz-game/quiz"
)

const filename = "questions.csv"

func main() {
	quiz, err := quiz.New(filename)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 5; i++ {
		que, err := quiz.GetRandom()
		if err != nil {
			log.Fatalln(err)
		}
		printQuestion(que)
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
