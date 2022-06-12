package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	answer   string
}

func main() {

	var file_name string
	var time_available int
	var shuffle bool

	flag.BoolVar(&shuffle, "sh", false, "whether to shuffle the question list extracted from the CSV file or not")
	flag.StringVar(&file_name, "f", "problems.csv", "the CSV file containing the questions and answers (separator : ',')")
	flag.IntVar(&time_available, "t", 30, "timelimit for the game in seconds")
	flag.Parse()

	file, err := os.Open(file_name)
	check(err)

	question_array := questions(file)
	if shuffle {
		println("Shuffling")
		shuffle_slice(question_array)
	}
	answer_pipeline := make(chan bool)

	// start with the user input
	fmt.Printf("Enter any key to start")
	cmd_reader := bufio.NewReader(os.Stdin)
	_, err = cmd_reader.ReadString('\n')
	check(err)

	go ask_questions(cmd_reader, question_array, answer_pipeline)
	timeout := time.After(time.Duration(time_available) * time.Second)

	correct_answers := 0
GAME:
	for {
		select {
		case a, open := <-answer_pipeline:
			if a {
				correct_answers++
			} else if !open {
				break GAME
			}
		case <-timeout:
			fmt.Printf("\nTime is up!\n")
			break GAME
		}
	}

	fmt.Printf("Game is over, your score is %v/%v", correct_answers, len(question_array))
}

func questions(f *os.File) (questions []Question) {
	csv_reader := csv.NewReader(f)
	records, err := csv_reader.ReadAll() // I didn't like this
	check(err)

	for _, r := range records {
		questions = append(questions, Question{question: normalize_string(r[0]), answer: normalize_string(r[1])})
	}

	return

}

func ask_questions(cmd_reader *bufio.Reader, question_array []Question, answer_pipeline chan<- bool) {
	defer close(answer_pipeline)
	var input, user_answer string
	var err error
	for _, q := range question_array {
		fmt.Printf("%v : ", q.question)
		input, err = cmd_reader.ReadString('\n')
		check(err)
		user_answer = strings.TrimSuffix(input, "\n")
		user_answer = normalize_string(user_answer)

		answer_pipeline <- user_answer == q.answer
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize_string(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func shuffle_slice[T any](slice []T) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
