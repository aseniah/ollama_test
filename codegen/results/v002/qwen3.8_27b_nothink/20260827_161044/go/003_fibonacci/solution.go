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
	
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return
	}
	
	if n < 1 {
		return
	}
	
	// Generate Fibonacci numbers up to and including the largest that doesn't exceed N
	// Start with 1, 1, 2, 3, 5, ...
	a := 1
	b := 1
	
	if n >= 1 {
		fmt.Println(a)
	}
	
	if n >= 1 {
		fmt.Println(b)
	}
	
	// Now generate subsequent Fibonacci numbers
	prev1 := 1
	prev2 := 1
	for {
		next := prev1 + prev2
		if next > n {
			break
		}
		fmt.Println(next)
		prev1 = prev2
		prev2 = next
	}
}