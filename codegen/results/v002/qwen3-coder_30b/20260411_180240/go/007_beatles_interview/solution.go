package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	Birthday   string     `json:"Birthday"`
	Age        int        `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

type Relative struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Skip header
	records = records[1:]

	var people []Person

	for _, record := range records {
		// Parse name
		nameParts := strings.Split(record[0], " ")
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday
		birthday := parseDate(record[1])

		// Calculate age as of July 1, 2025
		age := calculateAge(birthday, "2025-07-01")

		// Parse relatives
		relatives := make([]Relative, 0)

		// Father
		if record[3] != "null" {
			fatherName := strings.Split(record[3], " ")
			relatives = append(relatives, Relative{
				FirstName:    fatherName[0],
				LastName:     fatherName[len(fatherName)-1],
				Relationship: "Father",
			})
		}

		// Mother
		if record[4] != "null" {
			motherName := strings.Split(record[4], " ")
			relatives = append(relatives, Relative{
				FirstName:    motherName[0],
				LastName:     motherName[len(motherName)-1],
				Relationship: "Mother",
			})
		}

		// Brother
		if record[5] != "null" {
			brotherName := strings.Split(record[5], " ")
			relatives = append(relatives, Relative{
				FirstName:    brotherName[0],
				LastName:     brotherName[len(brotherName)-1],
				Relationship: "Brother",
			})
		}

		// Sister
		if record[6] != "null" {
			sisterName := strings.Split(record[6], " ")
			relatives = append(relatives, Relative{
				FirstName:    sisterName[0],
				LastName:     sisterName[len(sisterName)-1],
				Relationship: "Sister",
			})
		}

		person := Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   birthday,
			Age:        age,
			Relatives:  relatives,
		}

		people = append(people, person)
	}

	output, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

func parseDate(dateStr string) string {
	// Convert MM/DD/YYYY to YYYY-MM-DD
	parts := strings.Split(dateStr, "/")
	month := parts[0]
	day := parts[1]
	year := parts[2]

	// Ensure month and day are zero-padded
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}

	return fmt.Sprintf("%s-%s-%s", year, month, day)
}

func calculateAge(birthday string, referenceDate string) int {
	birth, _ := time.Parse("2006-01-02", birthday)
	ref, _ := time.Parse("2006-01-02", referenceDate)

	age := ref.Year() - birth.Year()
	if ref.Month() < birth.Month() || (ref.Month() == birth.Month() && ref.Day() < birth.Day()) {
		age--
	}

	return age
}