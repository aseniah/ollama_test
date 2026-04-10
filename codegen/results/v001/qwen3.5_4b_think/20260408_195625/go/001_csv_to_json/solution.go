package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string `json:"name"`
	Ages  int    `json:"age"`
	Email string `json:"email"`
	Score float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(rows) <= 1 {
		fmt.Println("[]")
		return
	}

	var students []Student
	for i, row := range rows {
		if i == 0 { // Skip header row
			continue
		}
		if len(row) < 4 {
			fmt.Fprintf(os.Stderr, "Invalid row %d: expected 4 columns\n", i+1)
			continue
		}
		
		name := strings.TrimSpace(row[0])
		age, err := strconv.Atoi(strings.TrimSpace(row[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid age in row %d: %v\n", i+1, err)
			continue
		}
		email := strings.TrimSpace(row[2])
		score, err := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid score in row %d: %v\n", i+1, err)
			continue
		}

		students = append(students, Student{
			Name:  name,
			Ages:  age,
			Email: email,
			Score: score,
		})
	}

	output, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}