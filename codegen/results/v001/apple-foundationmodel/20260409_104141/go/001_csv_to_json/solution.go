package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadCSV reads data from a CSV file and returns it as a slice of []struct{}
func ReadCSV(filePath string) ([]struct{}, error) {
	var data []struct{
		Name string
		Age  int
		Email string
		Score float64
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read header

	for scanner.Scan() {
		var name, age, email, score string
		err = scanner.ScanName() // Read Name
		if err != nil {
			return nil, fmt.Errorf("failed to parse name: %w", err)
		}
		err = scanner.ScanInt(&age) // Read Age
		if err != nil {
			return nil, fmt.Errorf("failed to parse age: %w", err)
		}
		err = scanner.ScanString(&email) // Read Email
		if err != nil {
			return nil, fmt.Errorf("failed to parse email: %w", err)
		}
		err = scanner.ScanFloat64(&score) // Read Score
		if err != nil {
			return nil, fmt.Errorf("failed to parse score: %w", err)
		}

		// Create a struct and append to slice
		data = append(data, struct {
			Name string
			Age  int
			Email string
			Score float64
		}{
			Name: name,
			Age:  age,
			Email: email,
			Score: score,
		})
	}

	return data, nil
}

func main() {
	filePath := "input/data.csv"
	data, err := ReadCSV(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Sort data by age for consistent output
	data = sortByAge(data)

	// Prepare JSON output
	jsonOutput := []byte{}
	jsonWriter := json.NewWriter(&jsonOutput)
	defer jsonWriter.Flush()

	if err := jsonWriter.Write(data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonOutput))
}

// sortByAge sorts the data array by age
func sortByAge(data []struct{}) []struct{} {
	return append([]struct{}{}, sort.Sort(data, func(i, j int) bool {
		return data[i].Age < data[j].Age
	})...)
}