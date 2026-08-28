package main

import "fmt"

func main() {
	// Check if command line arguments exist
	if len(fmt.Args()) == 0 {
		return // No argument provided, do nothing (though typical program behavior might exit)
	}

	var N int64
	fmt.Sscanf(fmt.Args()[0], "%d", &N)

	// If N < 1, print nothing and exit
	if N < 1 {
		return
	}

	f1 := int64(1) // First Fibonacci number
	f2 := int64(1) // Second Fibonacci number (also 1)
	hasPrinted := false

	// Print the first '1' if N >= 1
	if N >= 1 {
		fmt.Println(f1)
		hasPrinted = true
	}

	// Loop to find Fibonacci numbers <= N
	// We start by calculating the next one (1 + 1 = 2)
	for f2 <= N {
		next := f1 + f2
		f1 = f2
		f2 = next

		// If the new f2 is <= N, print it
		if f2 <= N {
			fmt.Println(f2)
		}
	}
}