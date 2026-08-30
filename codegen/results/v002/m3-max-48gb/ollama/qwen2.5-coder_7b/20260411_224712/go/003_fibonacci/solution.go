package main

import (
	"fmt"
	"os"
)

func generateFibonacci(max int) []int {
	if max < 1 {
		return []int{}
	}

	fib := []int{1, 1}
	for next := fib[len(fib)-1] + fib[len(fib)-2]; next <= max; {
		fib = append(fib, next)
		next = fib[len(fib)-1] + fib[len(fib)-2]
	}

	return fib
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fibonacci.go <integer>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid argument: %v\n", err)
		os.Exit(1)
	}

	fibNumbers := generateFibonacci(n)
	for _, num := range fibNumbers {
		fmt.Println(num)
	}
}