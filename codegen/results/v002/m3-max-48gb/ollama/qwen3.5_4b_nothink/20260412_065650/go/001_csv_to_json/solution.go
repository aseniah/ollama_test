package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Skip header row
	for i, record := range records {
		if i == 0 {
			continue
		}

		name := record[0]
		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid age for %s: %v\n", name, err)
			os.Exit(1)
		}
		email := record[2]
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid score for %s: %v\n", name, err)
			os.Exit(1)
		}

		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}

		people := make([]Person, i)
		copy(people, []Person{person})

		output, err := json.MarshalIndent(people, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(output))
		return
	}
}