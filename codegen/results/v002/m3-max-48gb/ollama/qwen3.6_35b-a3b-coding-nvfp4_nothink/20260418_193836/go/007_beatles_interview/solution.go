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
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	Birthday   string     `json:"Birthday"`
	Age        int        `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func parseBirthday(dateStr string) (time.Time, error) {
	// Format is M/D/YYYY or MM/DD/YYYY
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}
	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, err
	}
	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, err
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func calculateAge(birthday time.Time, referenceDate time.Time) int {
	age := referenceDate.Year() - birthday.Year()
	// Check if the birthday has occurred in the reference year
	if referenceDate.Year() == birthday.Year() {
		return 0
	}
	// Check month and day
	refDateInBirthdayYear := time.Date(referenceDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
	if referenceDate.Before(refDateInBirthdayYear) {
		age--
	}
	return age
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
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

	// Reference date: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading record: %v\n", err)
			continue
		}

		// Parse the record
		// Fields: Name, Birthday, Died, Father, Mother, Brother, Sister
		name := strings.TrimSpace(record[0])
		birthdayStr := strings.TrimSpace(record[1])
		// died is not needed for output
		fatherName := strings.TrimSpace(record[3])
		motherName := strings.TrimSpace(record[4])
		brotherName := strings.TrimSpace(record[5])
		sisterName := strings.TrimSpace(record[6])

		// Parse name into first and last
		nameParts := strings.Fields(name)
		var firstName, lastName string
		if len(nameParts) == 0 {
			continue
		} else if len(nameParts) == 1 {
			firstName = nameParts[0]
			lastName = ""
		} else {
			firstName = nameParts[0]
			lastName = nameParts[len(nameParts)-1]
		}

		// Parse birthday
		birthday, err := parseBirthday(birthdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday for %s: %v\n", name, err)
			continue
		}

		// Calculate age
		age := calculateAge(birthday, refDate)

		// Format birthday as YYYY-MM-DD
		birthdayStrFormatted := birthday.Format("2006-01-02")

		// Build relatives list
		var relatives []Relative

		if fatherName != "" && fatherName != "null" {
			fNameParts := strings.Fields(fatherName)
			fFirstName := fNameParts[0]
			fLastName := ""
			if len(fNameParts) > 1 {
				fLastName = fNameParts[len(fNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    fFirstName,
				LastName:     fLastName,
				Relationship: "Father",
			})
		}

		if motherName != "" && motherName != "null" {
			mNameParts := strings.Fields(motherName)
			mFirstName := mNameParts[0]
			mLastName := ""
			if len(mNameParts) > 1 {
				mLastName = mNameParts[len(mNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    mFirstName,
				LastName:     mLastName,
				Relationship: "Mother",
			})
		}

		if brotherName != "" && brotherName != "null" {
			bNameParts := strings.Fields(brotherName)
			bFirstName := bNameParts[0]
			bLastName := ""
			if len(bNameParts) > 1 {
				bLastName = bNameParts[len(bNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    bFirstName,
				LastName:     bLastName,
				Relationship: "Brother",
			})
		}

		if sisterName != "" && sisterName != "null" {
			sNameParts := strings.Fields(sisterName)
			sFirstName := sNameParts[0]
			sLastName := ""
			if len(sNameParts) > 1 {
				sLastName = sNameParts[len(sNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    sFirstName,
				LastName:     sLastName,
				Relationship: "Sister",
			})
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStrFormatted,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	// Output as JSON array
	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}