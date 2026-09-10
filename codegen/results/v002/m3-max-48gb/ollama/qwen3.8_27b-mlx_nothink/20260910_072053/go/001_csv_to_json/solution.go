package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	people := make([]Person, 0, len(records)-1)
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}
		if len(record) < 4 {
			continue
		}
		age, err := strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		score, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		people = append(people, Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}