package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Function to check if a number is prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Function to print all primes up to and including N
func printPrimes(N int) {
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // Remove the newline character

	N, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		return
	}

	printPrimes(N)
}