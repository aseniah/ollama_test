package main

import (
	"fmt"
	"os"
)

func fibonacci(n int) []int {
	if n < 1 {
		return []int{}
	}

	fib := make([]int, 0, 2)
	fib[0] = 1
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1] + fib[i-2])
	}
	return fib
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fibonacci.go <N>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Println("Invalid input. Please provide a positive integer.")
		return
	}

	fibonacciNumbers := fibonacci(n)

	for _, num := range fibonacciNumbers {
		if num > n {
			break
		}
		fmt.Print(num)
	}
}