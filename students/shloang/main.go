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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press enter to start the timer")
	scanner.Scan()
	//starting timer
	timer := time.NewTimer(3 * time.Second)
	inputChan := make(chan string)
	//scan user input, push to inputChan
	go func() {
		for _, item := range q {
			fmt.Printf(item.question + " is ")
			scanner.Scan()
			inputChan <- scanner.Text()
		}
	}()
	//handle input from inputChan or timeout
	tracker := 0
Iteraton:
	for i := 0; i < len(q); {
		select {
		case <-timer.C:
			fmt.Printf("\nTimeout\n")
			break Iteraton
		case answer := <-inputChan:
			if answer == q[i].answer {
				tracker++
			}
			i++
		}
	}
	fmt.Printf("%d right out of %d total", tracker, len(q))
}
