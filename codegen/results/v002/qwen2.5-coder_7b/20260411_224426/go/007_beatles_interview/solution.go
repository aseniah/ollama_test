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
	Relatives []Relative `json:"Relatives"`
}

type Relative struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var people []Person

	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		nameParts := strings.Split(record[0], " ")
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			lastName = nameParts[1]
		}

		birthday, _ := time.Parse("1/2/2006", record[1])
		age := calculateAge(birthday, 7, 1, 2025)

		var relatives []Relative
		addRelative(&relatives, "Father", record[3], lastName)
		addRelative(&relatives, "Mother", record[4], lastName)
		addRelative(&relatives, "Brother", record[5], lastName)
		addRelative(&relatives, "Sister", record[6], lastName)

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	jsonData, _ := json.MarshalIndent(people, "", "  ")
	fmt.Println(string(jsonData))
}

func calculateAge(birthday time.Time, month int, day int, year int) int {
	currentYear := year
	birthMonth := birthday.Month()
	birthDay := birthday.Day()

	if currentYear < birthday.Year() ||
		(currentYear == birthday.Year() && (birthMonth > month || (birthMonth == month && birthDay > day))) {
		currentYear--
	}

	return currentYear - birthday.Year()
}

func addRelative(relatives *[]Relative, relationship string, firstName string, lastName string) {
	if firstName != "" && lastName != "" {
		*relatives = append(*relatives, Relative{
			FirstName: firstName,
			LastName:  lastName,
			Relationship: relationship,
		})
	}
}