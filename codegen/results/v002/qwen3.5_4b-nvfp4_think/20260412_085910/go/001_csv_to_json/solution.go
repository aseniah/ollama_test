package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
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
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header
	var persons []Person
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
            os.Exit(1)
        }
        person := Person{
            Name:  row[0],
            Age:   parseInt(row[1]),
            Email: row[2],
            Score: parseFloat(row[3]),
        }
        persons = append(persons, person)
    }
	jsonBytes, err := json.Marshal(persons)
	if err != nil {
		os.Exit(1)
	}
	fmt.Print(string(jsonBytes))
}

func parseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func parseFloat(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}