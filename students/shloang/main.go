package main

import (
	"encoding/csv"
	"errors"
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

func cliCommunication() {}
