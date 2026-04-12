/**
 * This program takes a single integer N from the command line 
 * and prints all prime numbers up to and including N, one per line.
 * It uses the Sieve of Eratosthenes algorithm for efficiency.
 */

import process from 'node:process';

function solve(): void {
    // Get the command line argument
    const arg = process.argv[2];

    // Validate that an argument is provided
    if (arg === undefined) {
        return;
    }

    // Parse the argument to an integer
    const n = parseInt(arg, 10);

    // If N is not a valid number or less than 2, there are no primes to print
    if (isNaN(n) || n < 2) {
        return;
    }

    // Use Sieve of Eratosthenes to find primes up to N
    // Create a boolean array "isPrime[0..n]" and initialize
    // all entries it as true. A value in isPrime[i] will
    // finally be false if i is Not a prime, else true.
    const isPrime: boolean[] = new Array(n + 1).fill(true);
    isPrime[0] = false;
    isPrime[1] = false;

    for (let p = 2; p * p <= n; p++) {
        // If isPrime[p] is not changed, then it is a prime
        if (isPrime[p]) {
            // Update all multiples of p starting from p*p
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = false;
            }
        }
    }

    // Collect and print all primes
    let output = '';
    for (let p = 2; p <= n; p++) {
        if (isPrime[p]) {
            output += p + '\n';
        }
    }

    // Write to stdout
    process.stdout.write(output);
}

// Execute the solver
solve();