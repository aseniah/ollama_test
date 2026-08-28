package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Reference date for age calculation
var refDate = time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Birthday    string    `json:"Birthday"`
	Age         int       `json:"Age"`
	Relatives   []Relative `json:"Relatives"`
}

func parseBirthday(dateStr string) (time.Time, error) {
	// Handle format: M/D/YYYY or MM/D/YYYY or M/DD/YYYY etc.
	// The example input uses 10/9/1940 format.
	pattern := `^(\d{1,2})/(\d{1,2})/(\d{4})$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(dateStr)
	if len(matches) != 4 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}
	month, _ := strconv.Atoi(matches[1])
	day, _ := strconv.Atoi(matches[2])
	year, _ := strconv.Atoi(matches[3])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func calculateAge(birthday time.Time, ref time.Time) int {
	age := ref.Year() - birthday.Year()
	if time.Date(ref.Year(), ref.Month(), ref.Day(), 0, 0, 0, 0, time.UTC).Before(time.Date(birthday.Year(), ref.Month(), ref.Day(), 0, 0, 0, 0, time.UTC)) {
		age--
	}
	return age
}

func extractName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read header
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Map headers to indices
	nameIdx := -1
	birthdayIdx := -1
	fatherIdx := -1
	motherIdx := -1
	brotherIdx := -1
	sisterIdx := -1
	diedIdx := -1 // Not used for output but present in CSV

	for i, h := range header {
		switch strings.TrimSpace(h) {
		case "Name":
			nameIdx = i
		case "Birthday":
			birthdayIdx = i
		case "Father":
			fatherIdx = i
		case "Mother":
			motherIdx = i
		case "Brother":
			brotherIdx = i
		case "Sister":
			sisterIdx = i
		case "Died":
			diedIdx = i
		}
	}

	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV row: %v\n", err)
			os.Exit(1)
		}

		fullName := record[nameIdx]
		birthdayStr := record[birthdayIdx]

		firstName, lastName := extractName(fullName)

		bday, err := parseBirthday(birthdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday for %s: %v\n", fullName, err)
			os.Exit(1)
		}

		age := calculateAge(bday, refDate)

		var relatives []Relative

		// Father
		fatherStr := strings.TrimSpace(record[fatherIdx])
		if fatherStr != "null" && fatherStr != "" {
			fName, fNameLast := extractName(fatherStr)
			relatives = append(relatives, Relative{FirstName: fName, LastName: fNameLast, Relationship: "Father"})
		}

		// Mother
		motherStr := strings.TrimSpace(record[motherIdx])
		if motherStr != "null" && motherStr != "" {
			mName, mNameLast := extractName(motherStr)
			relatives = append(relatives, Relative{FirstName: mName, LastName: mNameLast, Relationship: "Mother"})
		}

		// Brother
		brotherStr := strings.TrimSpace(record[brotherIdx])
		if brotherStr != "null" && brotherStr != "" {
			bName, bNameLast := extractName(brotherStr)
			relatives = append(relatives, Relative{FirstName: bName, LastName: bNameLast, Relationship: "Brother"})
		}

		// Sister
		sisterStr := strings.TrimSpace(record[sisterIdx])
		if sisterStr != "null" && sisterStr != "" {
			sName, sNameLast := extractName(sisterStr)
			relatives = append(relatives, Relative{FirstName: sName, LastName: sNameLast, Relationship: "Sister"})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bday.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}