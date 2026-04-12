package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Person struct {
	Name  string
	Age   int
	Email string
	Score float64
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

	var people []Person

	for i, record := range records {
		// Skip header row
		if i == 0 {
			continue
		}
		// Ensure record has enough fields
		if len(record) < 4 {
			continue
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		people = append(people, Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		})
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}