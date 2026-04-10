package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("N", 0, "The maximum Fibonacci value")
	flag.Parse()

	if *n < 1 {
		return
	}

	// First two Fibonacci numbers
	a, b := 1, 1
	
	// Print first number
	fmt.Println(a)
	
	// If n is 1, we're done after the first number
	if *n == 1 {
		return
	}
	
	// Print second number (always 1 when N >= 1)
	fmt.Println(b)
	
	// Continue generating Fibonacci numbers
	for b <= *n {
		c := a + b
		a = b
		b = c
		
		if c <= *n {
			fmt.Println(c)
		} else {
			break
		}
	}
}