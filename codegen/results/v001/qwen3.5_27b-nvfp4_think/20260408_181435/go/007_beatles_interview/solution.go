package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	BirthDate  string `json:"birth_date"`
	Age        int    `json:"age"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
}

func calculateAge(birthDateStr string, referenceDate time.Time) int {
	if birthDateStr == "" {
		return 0
	}

	var birthDate time.Time
	var err error

	for _, format := range []string{
		"2006-01-02",
		"01/02/2006",
		"02/01/2006",
		"Jan 2, 2006",
		"January 2, 2006",
	} {
		birthDate, err = time.Parse(format, birthDateStr)
		if err == nil {
			break
		}
	}

	if err != nil {
		return 0
	}

	age := referenceDate.Year() - birthDate.Year()
	if !birthDate.After(referenceDate.AddDate(-age, 0, 0)) {
		age--
	}

	return age
}

func mapCSVToPerson(record []string, headers []string) (Person, error) {
	p := Person{}

	indexMap := make(map[string]int)
	for i, header := range headers {
		header = strings.TrimSpace(header)
		lowerHeader := strings.ToLower(header)
		switch lowerHeader {
		case "id":
			indexMap["id"] = i
		case "first name", "firstname", "firstName":
			indexMap["first_name"] = i
		case "last name", "lastname", "lastName":
			indexMap["last_name"] = i
		case "birth date", "birthdate", "dob", "date of birth", "birthday":
			indexMap["birth_date"] = i
		case "email":
			indexMap["email"] = i
		case "phone", "telephone", "tel":
			indexMap["phone"] = i
		case "city":
			indexMap["city"] = i
		case "state":
			indexMap["state"] = i
		case "country":
			indexMap["country"] = i
		}
	}

	getValue := func(key string) string {
		if idx, ok := indexMap[key]; ok && idx < len(record) {
			return strings.TrimSpace(record[idx])
		}
		return ""
	}

	p.ID = getValue("id")
	p.FirstName = getValue("first_name")
	p.LastName = getValue("last_name")
	p.BirthDate = getValue("birth_date")
	p.Email = getValue("email")
	p.Phone = getValue("phone")
	p.City = getValue("city")
	p.State = getValue("state")
	p.Country = getValue("country")

	// Calculate age as of July 1, 2025
	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	p.Age = calculateAge(p.BirthDate, referenceDate)

	return p, nil
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(allRecords) == 0 {
		jsonArray := make([]Person, 0)
		output, _ := json.MarshalIndent(jsonArray, "", "  ")
		fmt.Println(string(output))
		return
	}

	headers := allRecords[0]
	records := allRecords[1:]

	var people []Person
	for _, record := range records {
		if len(record) == 0 || (len(record) == 1 && record[0] == "") {
			continue
		}

		person, err := mapCSVToPerson(record, headers)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error mapping record: %v\n", err)
			continue
		}

		if person.ID == "" && person.FirstName == "" && person.LastName == "" {
			continue
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}