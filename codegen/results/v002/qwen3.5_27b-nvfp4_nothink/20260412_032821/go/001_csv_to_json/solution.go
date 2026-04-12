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

	var people []Person
	// Skip header (index 0)
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing Age for row", i+1, ":", err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing Score for row", i+1, ":", err)
			os.Exit(1)
		}

		person := Person{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}