package main

import (
	"fmt"
	"log"

	"github.com/AlinBorcea/quiz-game/quiz"
)

func main() {
	quiz, err := quiz.New("questions.csv")
	if err != nil {
		log.Fatalln(err)
	}

	question, err := quiz.GetQuestion(3)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(question)
}
