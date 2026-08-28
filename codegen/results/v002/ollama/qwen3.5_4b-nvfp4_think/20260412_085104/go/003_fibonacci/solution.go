package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <number>\n", os.Args[0])
		os.Exit(1)
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number format\n")
		os.Exit(1)
	}

	if N < 0 {
		return
	}

	var prev, curr int64 = 1, 1
	
	for curr <= int64(N) {
		fmt.Println(curr)
		next := prev + curr
		prev = curr
		curr = next
	}
}