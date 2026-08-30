package main

import (
	"fmt"
	"os"
)

// isPrime checks if a number n > 1 is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <N>\n", os.Args[0])
		os.Exit(1)
	}

	n, err := parseInt(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing N: %v\n", err)
		os.Exit(1)
	}

	// We need to find primes up to n. If n < 2, there are no primes.
	if n < 2 {
		return // No output
	}

	// Since we need to iterate through all numbers up to N, 
	// a simple loop is fine for the general case, though less efficient than Sieve of Eratosthenes.
	// However, the requirement is just "up to and including N". 
	// For very large N (e.g., > 10^9), this will TLE in pure Go without optimization, but that's 
	// acceptable for a basic prompt unless constraints are specified.
	// To be safe for typical interview/CP bounds (N <= 2*10^6 or similar), we can use Sieve.
	// But since the prompt doesn't specify N's range and "one per line" suggests output volume, 
	 // let's assume a reasonable upper bound where O(N log log N) might be needed if N is large.
	 // However, implementing a sieve in Go for arbitrary N (if user gives 10^9) is tricky due to memory.
	 // Let's stick with O(N sqrt(N)) per number check but optimized? 
	 // Actually, standard sieve is better. Let's implement Sieve of Eratosthenes.

	if n > 2*1000000 {
		// If N is extremely large (e.g., > 2 million), sieving all numbers up to N 
		// might consume too much memory/time for a simple script, but it's the correct approach for "all primes up to N".
		// We'll implement a sieve anyway.
	}

	// Create sieve array for numbers up to n.
	// Use byte slice for memory efficiency.
	sieve := make([]bool, n+1)
	// All entries are true (potentially prime) by default; but we need to set to false as non-prime.
	// Let's assume sieve[i] = true if i is prime. Initialize all to true, then mark 0 and 1 as false.
	for i := 2; i <= n; i++ {
		sieve[i] = true // initially we consider all > 1 as potentially prime
	}

	if n >= 2 {
		sieve[0] = false // not prime
		sieve[1] = false // not prime
	}

	// Sieve algorithm
	for p := 2; p <= n; p++ {
		if sieve[p] { // if p is prime
			if p*p <= n { // if p*p <= n, then mark multiples of p
				for i := p * p; i <= n; i += p {
					sieve[i] = false
				}
			}
		}
	}

	var count int
	// Count and store primes to output them one per line.
	// We collect them into a slice or channel? Actually, just loop and print.
	// However, if n is large, printing 10^9 numbers might exceed stdout limits in a single run (e.g., > 2GB). 
	 // But the problem says "print all prime numbers up to and including N". So we must do it.
	 // Let's print directly to avoid memory issues.
	for i := 2; i <= n; i++ {
		if sieve[i] {
			fmt.Println(i)
			count++
		}
	}

	if count == 0 {
		return // nothing printed, which satisfies "print nothing"
	}
}

// parseInt safely parses an int from a string argument.
func parseInt(s string) (int, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil || val < math.MinInt32 || val > math.MaxInt32 {
		return 0, fmt.Errorf("invalid integer: %s", s)
	}
	return int(val), nil
}