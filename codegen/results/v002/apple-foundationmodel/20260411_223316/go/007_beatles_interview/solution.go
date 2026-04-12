package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
	Age       int
	Relatives []map[string]string
}

func main() {
	// Read CSV file
	csvFile, err := os.Open("input/input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Read and process CSV file
	reader := bufio.NewReader(csvFile)
	var people []Person

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		fields := strings.Fields(line)
		if len(fields) < 5 {
			log.Fatalf("Malformed line: %s", line)
		}

		person := Person{
			FirstName: fields[0],
			LastName:  fields[1],
			Birthday:  fields[2],
			Age:      0,
		}

		// Extract birth year
		year, err := strconv.Atoi(strings.TrimSpace(fields[2]))
		if err != nil {
			log.Fatalf("Failed to parse birth year: %s", err)
		}

		// Calculate age
		person.Age = 2025 - year

		// Parse relatives
		for i := 3; i < len(fields); i++ {
			relative := map[string]string{}
			relativeKey := strings.TrimPrefix(fields[i], "{")
			relativeValue := strings.TrimSuffix(fields[i], "}")
			relative["FirstName"] = relativeKey
			relative["LastName"] = relativeValue
			person.Relatives = append(person.Relatives, relative)
		}

		people = append(people, person)
	}

	// Write JSON output
	jsonFile, err := os.Create("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Convert to JSON and write to file
	encoder := json.NewEncoder(jsonFile)
	if err := encoder.Encode(people); err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON data written to output.json")
}