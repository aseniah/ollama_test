package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName    string
	LastName     string
	Birthday     string
	Age          int
	Relatives    []map[string]string
	IsStillAlive bool
}

func parsePerson(row io.Reader, idx int) (*Person, error) {
	reader := bufio.NewReader(row)
	var firstName, lastName string
	var birthday, age int
	var relatives []map[string]string

	err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read line: %w", err)
	}

	err = strings.Split(strings.TrimSpace(err), ",")
	if err != nil {
		return nil, fmt.Errorf("failed to split comma-separated values: %w", err)
	}

	firstName, err = strconv.Atoi(err[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse age: %w", err)
	}
	lastName, err = strconv.Atoi(err[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse age: %w", err)
	}
	birthday, err = strconv.Atoi(err[2])
	if err != nil {
		return nil, fmt.Errorf("failed to parse age: %w", err)
	}
	age = firstName * 100 + lastName
	relatives = make([]map[string]string, 0)

	var relativeFirstName, relativeLastName, relationship string
	var relationshipIndex int
	for idx < len(err) {
		if err[idx] == ',' {
			relationshipIndex++
			continue
		}
		if relationshipIndex > 0 {
			relativeFirstName, err = strconv.Atoi(err[idx])
			if err != nil {
				return nil, fmt.Errorf("failed to parse age: %w", err)
			}
			relativeLastName, err = strconv.Atoi(err[idx+1])
			if err != nil {
				return nil, fmt.Errorf("failed to parse age: %w", err)
			}
			relationship = err[idx+2]
			relatives = append(relatives, map[string]string{
				"FirstName": relativeFirstName,
				"LastName": relativeLastName,
				"Relationship": relationship,
			})
			idx++
		} else {
			relativeFirstName, err = strconv.Atoi(err[idx])
			if err != nil {
				return nil, fmt.Errorf("failed to parse age: %w", err)
			}
			relativeLastName, err = strconv.Atoi(err[idx+1])
			if err != nil {
				return nil, fmt.Errorf("failed to parse age: %w", err)
			}
			relationship = err[idx+2]
			relatives = append(relatives, map[string]string{
				"FirstName": relativeFirstName,
				"LastName": relativeLastName,
				"Relationship": relationship,
			})
			idx++
		}
	}

	return &Person{
		FirstName:    firstName,
		LastName:     lastName,
		Birthday:     birthday,
		Age:          age,
		Relatives:    relatives,
		IsStillAlive: false,
	}, nil
}

func calculateAge(dob string, baseDate time.Time) int {
	dobTime, err := time.Parse("02/01/1900", dob)
	if err != nil {
		return 0
	}
	return int(time.Since(baseDate).Seconds()) / 365
}

func main() {
	if len(os.Args) < 2 || os.Args[1] != "csv" {
		log.Fatal("expected input file to be a CSV")
	}

	filePath := fmt.Sprintf("input/%s", os.Args[1])
	if _, err := os.Stat(filePath); err != nil {
		log.Fatalf("failed to stat file: %s", filePath)
	}

	// Read CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", filePath)
	}

	defer file.Close()

	var people []Person
	var idx int
	for {
		row, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read line: %s", err)
		}
		person, err := parsePerson(strings.NewReader(row), idx)
		if err != nil {
			log.Fatalf("failed to parse person: %s", err)
		}
		people = append(people, person)
		idx++
	}

	// Output JSON
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal JSON: %s", err)
	}

	fmt.Println(string(output))
}