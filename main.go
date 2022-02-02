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

	que, err := quiz.GetQuestion(2)
	if err != nil {
		log.Fatalln(err)
	}

	printQuestion(que)
}

func printQuestion(q *quiz.Question) {
	fmt.Println(q.Que)
	for i := 0; i < len(q.Answers)-1; i += 2 {
		fmt.Printf("%s\t%s\n", q.Answers[i], q.Answers[i+1])
	}
	fmt.Println(q.Correct)
}
