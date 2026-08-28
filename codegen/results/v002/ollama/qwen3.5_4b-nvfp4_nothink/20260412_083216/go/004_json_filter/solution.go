package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
)

type Record struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	// Read input file
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input file:", err)
		os.Exit(1)
	}

	// Parse JSON array
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		os.Exit(1)
	}

	var records []Record
	for _, rawItem := range raw {
		var r Record
		if err := json.Unmarshal(rawItem, &r); err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing record item:", err)
			os.Exit(1)
		}
		records = append(records, r)
	}

	// Filter: active is true and age >= 30
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort by name ascending
	// We need a sortable interface for strings in Go (standard library does not have String interface directly)
	// To avoid external dependencies like exp/slices or sort.Slice with custom comparators without import.
	// We implement slice sort manually.

	for i, j := 0, len(filtered)-1; i < j; i++ { // swap loop for bubble sort (simplified)
		if filtered[i].Name > filtered[j].Name {
			filtered[i], filtered[j] = filtered[j], filtered[i]
		} else if filtered[i].Name == filtered[j].Name && filtered[i].Age != filtered[j].Age {
			// If names are same, we could add secondary sort by age, but spec says name ascending.
			// For deterministic stability we might use index or random, but let's keep it simple.
			// Just swap if names equal to make stable? Not needed strictly.
			// Actually bubble sort does not guarantee full sorting for equality without tie-breaks.
			// Let's fix the logic: just compare and swap.
			fmt.Fprintln(os.Stderr, "Comparison issue in bubble sort") // debug output only
		}
	}

	// Since standard Go library doesn't have String interface for strings directly, we use sort.Slice (available since Go 1.20).
	// But to be self-contained, implement a simple insertion sort:

	for i := 1; i < len(filtered); i++ {
		key := filtered[i]
		j := i - 1
		for j >= 0 && key.Name > filtered[j].Name {
			j--
		}
		j++ // back up to insert position
		filtered = append(append(filter[:j], nil...), filter[j+1:]..., key)
	}

	// Convert to JSON and output
	out, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling output:", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}

func (string) sortOrder() constraints.Ordered { return 0 } // Dummy for compile-check only, not used