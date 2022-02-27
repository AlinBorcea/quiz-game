package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AlinBorcea/quiz-game/quiz"
	"github.com/AlinBorcea/quiz-game/user"
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

	reader := bufio.NewReader(os.Stdin)
	runLoop(q, reader)

	name, err := readString(reader)
	if err != nil {
		log.Fatalln(err)
	}

	if err = postResult(name, q.Result()); err != nil {
		log.Fatalln(err)
	}
}

func runLoop(q *quiz.Quiz, reader *bufio.Reader) {
	var option int = 1
	var err error = nil

	for option != 0 {
		switch option {
		case 0:
			continue
		case 1:
			printMenu()
		case 2:
			fmt.Println("leaderboard")
		case 3:
			runTestRandom(q, reader)
		default:
			fmt.Println("does not exist")
		}

		option, err = readInt(reader)
		if err != nil {
			break
		}
	}
}

func printMenu() {
	fmt.Println("0 -> exit")
	fmt.Println("1 -> menu")
	fmt.Println("2 -> leaderboard")
	fmt.Println("3 -> quiz")
}

func runTestRandom(q *quiz.Quiz, reader *bufio.Reader) {
	var ans int

	que, err := q.RandomQuestion()
	for err == nil {
		printQuestion(&que)
		ans, _ = readInt(reader)
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

func readInt(reader *bufio.Reader) (int, error) {
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

func readString(reader *bufio.Reader) (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")

	return str, nil
}

func postResult(name string, result int) error {
	u := user.User{Name: name, Score: result}
	return user.PostUserScore(&u, scoreFilename)
}
