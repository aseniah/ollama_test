#!/usr/bin/env tsx
import process from 'node:process';

/**
 * This program implements the Sieve of Eratosthenes algorithm to find 
 * all prime numbers up to and including a given integer N.
 * 
 * Complexity:
 * Time: O(N log log N)
 * Space: O(N)
 */
function solve(): void {
    // Read the argument from command line (process.argv[2])
    const input = process.argv[2];

    // If no argument is provided, the program exits silently.
    if (input === undefined) {
        return;
    }

    // Parse the input as an integer.
    const n = parseInt(input, 10);

    // If N is not a number, or N is less than 2, there are no primes to print.
    if (isNaN(n) || n < 2) {
        return;
    }

    /**
     * We use a Uint8Array to represent the primality of numbers.
     * Uint8Array is much more memory-efficient than a standard Boolean array
     * in Node.js/V8 for large values of N.
     * index 'i' contains 1 if 'i' is prime, and 0 otherwise.
     */
    const isPrime = new Uint8Array(n + 1);
    
    // Initialize the array: assume all numbers are prime initially.
    isPrime.fill(1);

    // 0 and 1 are by definition not prime numbers.
    isPrime[0] = 0;
    isPrime[1] = 0;

    // Sieve of Eratosthenes algorithm implementation.
    const limit = Math.sqrt(n);
    for (let p = 2; p <= limit; p++) {
        // If isPrime[p] is not changed, then it is a prime.
        if (isPrime[p] === 1) {
            // Update all multiples of p starting from p*p.
            // We start at p*p because smaller multiples would have been
            // marked by previous prime factors.
            for (let multiple = p * p; multiple <= n; multiple += p) {
                isPrime[multiple] = 0;
            }
        }
    }

    /**
     * To optimize output performance, we use a buffer to aggregate
     * prime numbers before writing to stdout. This reduces the number 
     * of system calls (syscalls) which can be a bottleneck for large N.
     */
    let buffer = "";
    const BUFFER_THRESHOLD = 1000; // Flush buffer every 1000 primes found.
    let primeCount = 0;

    for (let i = 2; i <= n; i++) {
        if (isPrime[i] === 1) {
            buffer += i + "\n";
            primeCount++;

            // If buffer reaches threshold, write to stdout and clear buffer.
            if (primeCount >= BUFFER_THRESHOLD) {
                process.stdout.write(buffer);
                buffer = "";
                primeCount = 0;
            }
        }
    }

    // Write any remaining content in the buffer to stdout.
    if (buffer.length > 0) {
        process.stdout.write(buffer);
    }
}

// Execute the solver.
solve();