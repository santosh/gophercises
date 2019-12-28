package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Problem struct works as a proxy for a CSV row.
type Problem struct {
	question string
	answer   string
}

// ParseCSV takes a two column csv one having question and other answer
// and cast into Problem struct
func ParseCSV(rows [][]string) []Problem {
	problems := make([]Problem, len(rows))
	for i, row := range rows {
		problems[i] = Problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}
	return problems
}

func main() {
	filenamePtr := flag.String("csv", "problems.csv", "a csv file with ques,ans")
	timeLimitPtr := flag.Int("tlimit", 30, "time limit after which quiz will end")
	flag.Parse()

	csvFile, err := os.Open(*filenamePtr)
	handleError(err)
	timer := time.NewTimer(time.Duration(*timeLimitPtr) * time.Second)

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	handleError(err)

	problems := ParseCSV(rows)
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerChannel := make(chan string)
		go func() {
			var response string
			fmt.Scanln(&response)
			answerChannel <- response
		}()
		select {
		case <-timer.C:
			fmt.Printf("\n%d answers correct of %d questions.\n", correct, len(problems))
			return
		case response := <-answerChannel:
			if response == problem.answer {
				correct++
			}
		}
	}
	fmt.Printf("\n%d answers correct of %d questions.\n", correct, len(problems))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
