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
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Birthday  string `json:"Birthday"`
	Age       int    `json:"Age"`
	Relatives []struct {
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
		Relationship string `json:"Relationship"`
	} `json:"Relatives"`
}

func main() {
	// Open the input CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header
	header, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the header to find the index of each column
	indexName := -1
	indexBirthday := -1
	indexDied := -1
	indexFather := -1
	indexMother := -1
	indexBrother := -1
	indexSister := -1

	for i, col := range header {
		switch strings.TrimSpace(col) {
		case "Name":
			indexName = i
		case "Birthday":
			indexBirthday = i
		case "Died":
			indexDied = i
		case "Father":
			indexFather = i
		case "Mother":
			indexMother = i
		case "Brother":
			indexBrother = i
		case "Sister":
			indexSister = i
		}
	}

	if indexName == -1 || indexBirthday == -1 || (indexDied != -1 && indexFather == -1) {
		fmt.Println("CSV header is missing required columns.")
		return
	}

	// Read the data rows
	var people []Person

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		firstName := strings.TrimSpace(row[indexName])
		birthdayStr := strings.TrimSpace(row[indexBirthday])

		parts := strings.SplitN(firstName, " ", 2)
		firstName = parts[0]
		lastName := ""
		if len(parts) > 1 {
			lastName = parts[1]
		}

		birthday, err := time.Parse("1/2/06", birthdayStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		died := false
		if indexDied != -1 {
			diedStr := strings.TrimSpace(row[indexDied])
			if diedStr != "" {
				died = true
			}
		}

		fatherFirstName := ""
		fatherLastName := ""
		if indexFather != -1 {
			fatherFullName := strings.TrimSpace(row[indexFather])
			parts := strings.SplitN(fatherFullName, " ", 2)
			fatherFirstName = parts[0]
			if len(parts) > 1 {
				fatherLastName = parts[1]
			}
		}

		motherFirstName := ""
		motherLastName := ""
		if indexMother != -1 {
			motherFullName := strings.TrimSpace(row[indexMother])
			parts := strings.SplitN(motherFullName, " ", 2)
			motherFirstName = parts[0]
			if len(parts) > 1 {
				motherLastName = parts[1]
			}
		}

		brotherFirstName := ""
		brotherLastName := ""
		if indexBrother != -1 && row[indexBrother] != "" {
			brotherFullName := strings.TrimSpace(row[indexBrother])
			parts := strings.SplitN(brotherFullName, " ", 2)
			brotherFirstName = parts[0]
			if len(parts) > 1 {
				brotherLastName = parts[1]
			}
		}

		sisterFirstName := ""
		sisterLastName := ""
		if indexSister != -1 && row[indexSister] != "" {
			sisterFullName := strings.TrimSpace(row[indexSister])
			parts := strings.SplitN(sisterFullName, " ", 2)
			sisterFirstName = parts[0]
			if len(parts) > 1 {
				sisterLastName = parts[1]
			}
		}

		// Calculate age as of July 1, 2025
		birthDate := time.Date(birthday.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
		targetDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

		var age int
		if birthDate.Before(targetDate) {
			age = targetDate.Year() - birthDate.Year()
			if birthDate.After(targetDate.AddDate(age, 0, 0)) {
				age--
			}
		}

		// Create the person object
		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday.Format("2006-01-02"),
			Age:       age,
			Relatives: []struct {
				FirstName string `json:"FirstName"`
				LastName  string `json:"LastName"`
				Relationship string `json:"Relationship"`
			}{
				{Father, fatherLastName, "Father"},
				{Mother, motherLastName, "Mother"},
				{Brother, brotherLastName, "Brother"},
				{Sister, sisterLastName, "Sister"},
			},
		}

		// Add the person to the list
		people = append(people, person)
	}

	// Print the JSON array to stdout
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}