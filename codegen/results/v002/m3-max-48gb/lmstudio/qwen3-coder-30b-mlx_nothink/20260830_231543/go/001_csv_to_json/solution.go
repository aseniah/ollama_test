package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
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
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	var people []Person
	
	// Read header row
	header, err := reader.Read()
	if err != nil {
		panic(err)
	}
	
	// Process each data row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		
		// Parse age
		age, err := strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			panic(err)
		}
		
		// Parse score
		score, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		if err != nil {
			panic(err)
		}
		
		person := Person{
			Name:  strings.TrimSpace(record[0]),
			Age:   age,
			Email: strings.TrimSpace(record[2]),
			Score: score,
		}
		
		people = append(people, person)
	}
	
	// Output JSON
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s", jsonBytes)
}