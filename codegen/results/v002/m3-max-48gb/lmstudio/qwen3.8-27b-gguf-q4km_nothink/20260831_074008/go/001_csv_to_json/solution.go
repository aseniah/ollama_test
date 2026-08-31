package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read and skip header
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}
	
	// We'll use the header to determine column indices
	// For simplicity, assume the columns are in the expected order: Name, Age, Email, Score
	// But let's be robust and use the header to map columns
	
	colIndex := make(map[string]int)
	for i, h := range header {
		colIndex[h] = i
	}
	
	nameIdx, _ := colIndex["Name"]
	ageIdx, _ := colIndex["Age"]
	emailIdx, _ := colIndex["Email"]
	scoreIdx, _ := colIndex["Score"]
	
	var records []Record
	
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading row: %v\n", err)
			break
		}
		
		if len(row) <= max(nameIdx, max(ageIdx, max(emailIdx, scoreIdx))) {
			continue
		}
		
		age, err := strconv.Atoi(row[ageIdx])
		if err != nil {
			age = 0
		}
		
		score, err := strconv.ParseFloat(row[scoreIdx], 64)
		if err != nil {
			score = 0.0
		}
		
		record := Record{
			Name:  row[nameIdx],
			Age:   age,
			Email: row[emailIdx],
			Score: score,
		}
		
		records = append(records, record)
	}
	
	output, err := json.Marshal(records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println(string(output))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}