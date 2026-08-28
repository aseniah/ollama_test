package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
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
	
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var results []Person
	
	for {
		record, err := reader.Read()
		if err == csv.ErrEOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading record: %v\n", err)
			os.Exit(1)
		}

		if len(record) != 4 {
			fmt.Fprintf(os.Stderr, "Record has unexpected number of fields: %d\n", len(record))
			continue
		}

		person := Person{
			Name:   record[0],
			Age:    parseInt(record[1]),
			Email:  record[2],
			Score:  parseFloat(record[3]),
		}
		results = append(results, person)
	}

	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	_, err = fmt.Fprint(os.Stdout, string(output))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to stdout: %v\n", err)
		os.Exit(1)
	}
}

func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Invalid integer: %s", s))
	}
	return val
}

// Note: We need to import "strconv" here.
// Correcting the imports section below for the final output.

func parseFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("Invalid float: %s", s))
	}
	return val
}