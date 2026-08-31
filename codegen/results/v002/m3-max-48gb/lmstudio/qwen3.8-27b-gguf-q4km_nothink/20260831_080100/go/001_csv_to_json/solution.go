package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		// No data rows, output empty array
		fmt.Println("[]")
		return
	}

	var people []Person
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}

		name := row[0]
		ageStr := row[1]
		email := row[2]
		scoreStr := row[3]

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			age = 0
		}

		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			score = 0.0
		}

		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}
		people = append(people, person)
	}

	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}