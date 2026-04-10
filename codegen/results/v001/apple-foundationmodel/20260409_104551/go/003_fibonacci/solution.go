package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the integer argument N from the command line
	if len(os.Args) < 2 {
		fmt.Println("Usage: fibonacci.go <N>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error reading argument:", err)
		return
	}

	// Print nothing if N is less than 1
	if n < 1 {
		return
	}

	// Initialize the first two Fibonacci numbers
	fib := []int{1, 1}

	// Print all Fibonacci numbers up to and including the largest that does not exceed N
	for i := 2; i < len(fib); i++ {
		nextFib := fib[i-1] + fib[i-2]
		if nextFib > n {
			break
		}
		fib = append(fib, nextFib)
	}

	for _, num := range fib {
		fmt.Print(num)
	}
}