package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Birthday  string   `json:"Birthday"`
	Age       int      `json:"Age"`
	Relatives []Family `json:"Relatives"`
}

type Family struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	// Open the input CSV file
	inputFile, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Create a new CSV reader
	reader := csv.NewReader(inputFile)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Skip the header row and create a slice of Person structs
	var people []Person
	for i, record := range records {
		if i == 0 {
			continue // skip header row
		}
		person := Person{
			FirstName: strings.Split(record[0], " ")[1],
			LastName:  strings.Split(record[0], " ")[2],
			Birthday:  formatBirthday(record[1]),
		}

		// Calculate age as of July 1, 2025
		birthDate := parseDate(person.Birthday)
		person.Age = calculateAge(birthDate, 2025)

		// Parse relatives and add to Person struct
		relatives := []Family{}
		if record[4] != "" {
			relatives = append(relatives, Family{FirstName: record[4], LastName: "", Relationship: "Father"})
		}
		if record[5] != "" {
			relatives = append(relatives, Family{FirstName: record[5], LastName: "", Relationship: "Mother"})
		}
		if record[6] != "" {
			relatives = append(relatives, Family{FirstName: record[6], LastName: "", Relationship: "Brother"})
		}
		if record[7] != "" {
			relatives = append(relatives, Family{FirstName: record[7], LastName: "", Relationship: "Sister"})
		}
		person.Relatives = relatives

		people = append(people, person)
	}

	// Print the output JSON array
	jsonData, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(jsonData))
}

func formatBirthday(birthday string) string {
	parts := strings.Split(birthday, "/")
	return parts[2] + "-" + parts[0] + "-" + parts[1]
}

func parseDate(date string) (int, int, int) {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	return year, month, day
}

func calculateAge(birthDate string, currentYear int) int {
	year, _, _ := parseDate(birthDate)
	return currentYear - year
}