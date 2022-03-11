package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
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
	timer := time.NewTimer(3 * time.Second)
	tracker := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press enter to start the timer")
	scanner.Scan()
Iteraton:
	for _, item := range q {
		select {
		case <-timer.C:
			fmt.Printf("timeout\n")
			break Iteraton
		default:
			fmt.Printf(item.question + " is ")
			scanner.Scan()
			answer := scanner.Text()
			if answer == item.answer {
				tracker++
			}
		}
	}
	fmt.Printf("%d out of %d right", tracker, len(q))
}
