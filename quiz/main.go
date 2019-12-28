package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func parseCSV(rows [][]string) []problem {
	problems := make([]problem, len(rows))
	for i, row := range rows {
		problems[i] = problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}
	return problems
}

func main() {
	filenamePtr := flag.String("csv", "problems.csv", "a csv file with ques,ans")
	// timeLimitPtr := flag.Int("tlimit", 30, "time limit after which quiz will end")
	flag.Parse()

	csvFile, err := os.Open(*filenamePtr)
	handleError(err)

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	handleError(err)

	problems := parseCSV(rows)
	correct := 0

	for _, problem := range problems {
		fmt.Printf("What is %s? ", problem.question)
		var response string
		fmt.Scanln(&response)

		if response == problem.answer {
			correct++
		}
	}

	fmt.Printf("%d answers correct out of %d questions.\n", correct, len(problems))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
