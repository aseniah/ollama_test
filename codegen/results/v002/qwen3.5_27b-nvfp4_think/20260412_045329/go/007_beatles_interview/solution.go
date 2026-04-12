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

	var people []Person
	calcDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
			os.Exit(1)
		}

		name := strings.TrimSpace(record[0])
		firstName, lastName := splitName(name)

		birthdayStr := strings.TrimSpace(record[1])
		birthday := parseDate(birthdayStr)

		birthdayFormatted := formatISODate(birthday)

		age := calculateAge(birthday, calcDate)

		var relatives []Relative

		father := strings.TrimSpace(record[3])
		if father != "null" && strings.TrimSpace(father) != "" {
			firstF, lastF := splitName(father)
			relatives = append(relatives, Relative{FirstName: firstF, LastName: lastF, Relationship: "Father"})
		}

		mother := strings.TrimSpace(record[4])
		if mother != "null" && strings.TrimSpace(mother) != "" {
			firstM, lastM := splitName(mother)
			relatives = append(relatives, Relative{FirstName: firstM, LastName: lastM, Relationship: "Mother"})
		}

		brother := strings.TrimSpace(record[5])
		if brother != "null" && strings.TrimSpace(brother) != "" {
			firstB, lastB := splitName(brother)
			relatives = append(relatives, Relative{FirstName: firstB, LastName: lastB, Relationship: "Brother"})
		}

		sister := strings.TrimSpace(record[6])
		if sister != "null" && strings.TrimSpace(sister) != "" {
			firstS, lastS := splitName(sister)
			relatives = append(relatives, Relative{FirstName: firstS, LastName: lastS, Relationship: "Sister"})
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	outputJSON, err := json.MarshalIndent(people, "", "   ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(outputJSON))
}

func splitName(fullName string) (string, string) {
	parts := strings.Split(fullName, " ")
	if len(parts) == 0 {
		return "", ""
	} else if len(parts) == 1 {
		return parts[0], ""
	}
	firstName := parts[0]
	lastName = strings.Join(parts[1:], " ")
	return firstName, lastName
}

func parseDate(dateStr string) time.Time {
	formats := []string{
		"1/2/2006",
		"1/2/06",
		"01/02/2006",
		"01/02/06",
	}

	for _, format := range formats {
		if parsed, err := time.Parse(format, dateStr); err == nil {
			return parsed
		}
	}
	return time.Time{}
}

func formatISODate(t time.Time) string {
	return t.Format("2006-01-02")
}

func calculateAge(birthday time.Time, calcDate time.Date) int {
	age := calcDate.Year() - birthday.Year()
	if !calcDate.Before(time.Date(calcDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)) {
		age--
	}
	return age
}

func isValidDate(dateStr string) bool {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "null" || dateStr == "" {
		return false
	}
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return false
	}
	for _, part := range parts {
		if _, err := strconv.Atoi(strings.TrimSpace(part)); err != nil {
			return false
		}
	}
	return true
}