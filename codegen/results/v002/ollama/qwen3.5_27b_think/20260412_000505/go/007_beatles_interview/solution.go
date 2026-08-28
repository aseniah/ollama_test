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

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		os.Exit(1)
	}

	var people []Person
	header := records[0]

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 7 {
			continue
		}

		name := record[0]
		birthday := record[1]
		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		p := parseRecord(name, birthday, father, mother, brother, sister)
		people = append(people, p)
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}

func parseRecord(name, birthdayStr, father, mother, brother, sister string) Person {
	nameParts := strings.Fields(name)
	firstName := nameParts[0]
	lastName := nameParts[len(nameParts)-1]

	birthday := convertBirthday(birthdayStr)
	age := calculateAge(birthday)

	relatives := buildRelatives(father, "Father", mother, "Mother", brother, "Brother", sister, "Sister")

	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
		Age:       age,
		Relatives: relatives,
	}
}

func convertBirthday(dateStr string) string {
	parsed, err := time.Parse("1/2/2006", dateStr)
	if err != nil {
		parsed, err = time.Parse("1/2/2006", dateStr)
		if err != nil {
			return ""
		}
	}
	return parsed.Format("2006-01-02")
}

func calculateAge(birthday string) int {
	refDate, _ := time.Parse("2006-01-02", "2025-07-01")
	birthDate, _ := time.Parse("2006-01-02", birthday)

	age := refDate.Year() - birthDate.Year()
	if refDate.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

func buildRelatives(father, fatherRel, mother, motherRel, brother, brotherRel, sister, sisterRel string) []Relative {
	var relatives []Relative

	if father != "" && father != "null" {
		firstName, lastName := splitName(father)
		relatives = append(relatives, Relative{FirstName: firstName, LastName: lastName, Relationship: fatherRel})
	}
	if mother != "" && mother != "null" {
		firstName, lastName := splitName(mother)
		relatives = append(relatives, Relative{FirstName: firstName, LastName: lastName, Relationship: motherRel})
	}
	if brother != "" && brother != "null" {
		firstName, lastName := splitName(brother)
		relatives = append(relatives, Relative{FirstName: firstName, LastName: lastName, Relationship: brotherRel})
	}
	if sister != "" && sister != "null" {
		firstName, lastName := splitName(sister)
		relatives = append(relatives, Relative{FirstName: firstName, LastName: lastName, Relationship: sisterRel})
	}

	return relatives
}

func splitName(name string) (string, string) {
	parts := strings.Fields(name)
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
}

func parseCSVDate(dateStr string) time.Time {
	formats := []string{"1/2/2006", "1/2/06", "1/2/2006 15:04:05"}
	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t
		}
	}
	return time.Now()
}

func _calculateAge(birthday string) int {
	birthDate := parseCSVDate(birthday)
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	age := refDate.Year() - birthDate.Year()
	if refDate.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

func _unused() {
	strconv.Itoa(0)
}