package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Person struct {
	Name  string
	Age   int
	Email string
	Score float64
}

func main() {
	file, err := os.Open("input/data