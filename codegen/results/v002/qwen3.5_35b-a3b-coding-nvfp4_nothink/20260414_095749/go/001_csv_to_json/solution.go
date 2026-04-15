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
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		return
	}

	if len(records) < 2 {
		fmt.Fprintln(os.Stderr, "No data rows found")
		return
	}

	// Skip header (index 0)
	rows := records[1:]

	var people []Person

	for _, record := range rows {
		if len(record) < 4 {
			continue
		}

		name := record[0]
		ageStr := record[1]
		email := record[2]
		scoreStr := record[3]

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			continue
		}

		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			continue
		}

		people = append(people, Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}