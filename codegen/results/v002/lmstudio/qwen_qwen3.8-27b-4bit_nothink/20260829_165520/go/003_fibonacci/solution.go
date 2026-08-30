package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	
	if n < 1 {
		return
	}
	
	// Print all Fibonacci numbers up to and including the largest one that does not exceed N
	// Sequence: 1, 1, 2, 3, 5, ...
	a, b := 1, 1
	
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}