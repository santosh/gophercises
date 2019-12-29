package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *filenamePtr))
	}
	timer := time.NewTimer(time.Duration(*timeLimitPtr) * time.Second)

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file: %s\n")
	}

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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
