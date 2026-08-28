package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	
	// Read header
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading header:", err)
		return
	}

	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading record:", err)
			return
		}

		if len(record) < 4 {
			continue
		}

		var person Person
		person.Name = record[0]
		
		// Parse Age
		_, err = fmt.Sscanf(record[1], "%d", &person.Age)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing age:", err)
			return
		}
		
		person.Email = record[2]
		
		// Parse Score
		_, err = fmt.Sscanf(record[3], "%f", &person.Score)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing score:", err)
			return
		}

		people = append(people, person)
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}