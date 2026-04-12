package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if an argument was provided
	if len(os.Args) < 2 {
		return
	}

	// Parse the command line argument N
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		// If N is not a valid integer, we exit silently or could handle error
		return
	}

	// If N < 1, print nothing as per instructions
	if n < 1 {
		return
	}

	// Initialize the first two numbers of the Fibonacci sequence
	a := 1
	b := 1

	// Print the first number
	fmt.Println(a)

	// Handle the case where N is exactly 1
	// The sequence starts 1, 1, 2... so if N=1, we should print 1 and 1.
	// However, the logic below handles the progression.
	
	// If N is 1, the sequence up to N is just 1, 1.
	// Let's use a loop that calculates the next number and checks against N.
	
	// We already printed the first '1'. Now we check for the second '1' and onwards.
	// Since the sequence is 1, 1, 2, 3... we start checking from the second 1.
	
	current := b
	if current <= n {
		fmt.Println(current)
	} else {
		return
	}

	// Loop to calculate subsequent Fibonacci numbers
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		// Update values for next iteration
		a = b
		b = next
	}
}