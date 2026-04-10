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
	
	// Read header
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading header:", err)
		os.Exit(1)
	}

	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading record:", err)
			os.Exit(1)
		}

		if len(record) < 4 {
			continue
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error converting age:", err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error converting score:", err)
			os.Exit(1)
		}

		person := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}

		people = append(people, person)
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}