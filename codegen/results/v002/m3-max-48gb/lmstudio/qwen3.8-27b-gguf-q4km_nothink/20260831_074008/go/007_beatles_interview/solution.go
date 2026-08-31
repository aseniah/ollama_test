package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  string     `json:"Birthday"`
	Age       int        `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		return parts[0], ""
	}
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	// If there are more than 2 parts, first is first name, last is last name
	return parts[0], parts[len(parts)-1]
}

func splitRelName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	// If there are more than 2 parts, first is first name, last is last name
	return parts[0], parts[len(parts)-1]
}

func calculateAge(birthday string) int {
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	birth, err := time.Parse("1/2/2006", birthday)
	if err != nil {
		// Try other format
		birth, err = time.Parse("01/02/2006", birthday)
		if err != nil {
			return 0
		}
	}
	age := refDate.Year() - birth.Year()
	if refDate.YearDay() < birth.YearDay() {
		age--
	}
	return age
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	header := records[0]
	// Map column names to indices
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[col] = i
	}

	var result []Person

	for _, record := range records[1:] {
		name := strings.TrimSpace(record[colIndex["Name"]])
		birthday := strings.TrimSpace(record[colIndex["Birthday"]])

		firstName, lastName := splitName(name)

		relatives := []Relative{}

		// Father
		fatherName := strings.TrimSpace(record[colIndex["Father"]])
		if fatherName != "" && fatherName != "null" {
			fName, lName := splitRelName(fatherName)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: "Father",
			})
		}

		// Mother
		motherName := strings.TrimSpace(record[colIndex["Mother"]])
		if motherName != "" && motherName != "null" {
			fName, lName := splitRelName(motherName)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: "Mother",
			})
		}

		// Brother
		brotherName := strings.TrimSpace(record[colIndex["Brother"]])
		if brotherName != "" && brotherName != "null" {
			fName, lName := splitRelName(brotherName)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: "Brother",
			})
		}

		// Sister
		sisterName := strings.TrimSpace(record[colIndex["Sister"]])
		if sisterName != "" && sisterName != "null" {
			fName, lName := splitRelName(sisterName)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: "Sister",
			})
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  time.Now().Year() - time.Now().Year() + 1940, // placeholder
			Age:       calculateAge(birthday),
			Relatives: relatives,
		}

		// Format birthday as YYYY-MM-DD
		birth, _ := time.Parse("1/2/2006", birthday)
		if birth.IsZero() {
			birth, _ = time.Parse("01/02/2006", birthday)
		}
		person.Birthday = birth.Format("2006-01-02")

		result = append(result, person)
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}