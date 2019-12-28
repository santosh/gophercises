package main

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestParseCSVLen(t *testing.T) {
	csvData := strings.NewReader(strings.Join([]string{
		`20+1,21`,
		`9+1,10`,
	}, "\n"))

	reader := csv.NewReader(csvData)
	rows, _ := reader.ReadAll()

	problems := ParseCSV(rows)

	if len(problems) != 2 {
		t.Errorf("ParseCSV failed, expected %v, got %v", 2, len(problems))
	}
}

func TestParseCSVMembership(t *testing.T) {
	csvData := strings.NewReader(strings.Join([]string{
		`20+1,21`,
		`9+1,10`,
	}, "\n"))

	reader := csv.NewReader(csvData)
	rows, _ := reader.ReadAll()

	problems := ParseCSV(rows)

	if problems[0].answer != "21" {
		t.Errorf("Answer didn't match, expected %v, got %v", "21", problems[0].answer)
	}

	if problems[0].question != "20+1" {
		t.Errorf("Question didn't match, expected %v, got %v", "20+1", problems[0].question)
	}
}

func TestParseCSVEmpty(t *testing.T) {
	csvData := strings.NewReader(strings.Join([]string{
		``,
	}, "\n"))

	reader := csv.NewReader(csvData)
	rows, _ := reader.ReadAll()

	problems := ParseCSV(rows)

	if len(problems) != 0 {
		t.Errorf("Empty csv should not return any Problem{}")
	}
}
