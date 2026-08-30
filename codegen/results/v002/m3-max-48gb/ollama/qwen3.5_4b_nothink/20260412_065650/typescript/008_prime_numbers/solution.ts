import * as fs from 'fs';
import * as path from 'path';

// Get command line arguments
const args = process.argv.slice(2);

// Check if an argument was provided
if (args.length === 0) {
    process.exit(1);
}

// Parse the integer argument
const n = parseInt(args[0], 10);

// Validate that the argument is a valid positive integer
if (isNaN(n) || n <= 0 || !Number.isFinite(n)) {
    process.exit(1);
}

// Function to generate primes up to n using the Sieve of Eratosthenes
function generatePrimes(limit: number): number[] {
    if (limit < 2) {
        return [];
    }

    // Create a boolean array "isPrime" of size limit + 1, initialized to true
    const isPrime = new Array<boolean>(limit + 1).fill(true);

    isPrime[0] = false;
    isPrime[1] = false;

    // Only need to check up to sqrt(limit)
    for (let p = 2; p * p <= limit; p++) {
        if (isPrime[p]) {
            // Mark multiples of p as non-prime
            for (let multiple = p * p; multiple <= limit; multiple += p) {
                isPrime[multiple] = false;
            }
        }
    }

    // Collect all prime numbers
    const primes: number[] = [];
    for (let i = 2; i <= limit; i++) {
        if (isPrime[i]) {
            primes.push(i);
        }
    }

    return primes;
}

// Generate and print primes
const primes = generatePrimes(n);

// Print each prime number on a new line to stdout
for (const prime of primes) {
    console.log(prime);
}