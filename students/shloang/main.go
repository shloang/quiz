package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

type quizNode struct {
	question string
	answer   string
}

func main() {
	quiz, err := readCsvIntoQuizStruct()
	errorCheck(err)
	cliComms(quiz)
}

func readCsvIntoQuizStruct() ([]quizNode, error) {
	file, err := os.Open("problems.csv")
	errorCheck(err)
	defer file.Close()
	rawCsv, err := csv.NewReader(file).ReadAll()
	errorCheck(err)
	quizNodes := []quizNode{}
	for _, line := range rawCsv {
		if len(line) != 2 {
			return quizNodes, errors.New("CSV file row length is not 2")
		}
		quizNodes = append(quizNodes, quizNode{line[0], line[1]})
	}
	return quizNodes, nil
}

func errorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func cliComms(q []quizNode) {
	answer := ""
	tracker := 0
	for _, item := range q {
		fmt.Printf(item.question)
		fmt.Scanf("%s", &answer)
		if answer == item.answer {
			tracker++
		}
	}
	fmt.Printf("%d", tracker)
}
