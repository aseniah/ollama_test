/**
 * This script takes a single integer N from the command line 
 * and prints all prime numbers up to and including N, one per line.
 * 
 * It uses the Sieve of Eratosthenes algorithm for efficiency.
 */

import process from 'node:process';

/**
 * Computes primes up to N using the Sieve of Eratosthenes.
 * @param n The upper limit (inclusive)
 * @returns An array of prime numbers
 */
function getPrimesUpTo(n: number): number[] {
    if (n < 2) {
        return [];
    }

    // Initialize a boolean array "isPrime[0..n]" and fill it with true.
    // A value in isPrime[i] will finally be false if i is Not a prime, else true.
    const isPrime = new Uint8Array(n + 1).fill(1);
    isPrime[0] = 0;
    isPrime[1] = 0;

    const limit = Math.sqrt(n);

    for (let p = 2; p <= limit; p++) {
        // If isPrime[p] is not changed, then it is a prime
        if (isPrime[p] === 1) {
            // Update all multiples of p starting from p*p
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    const primes: number[] = [];
    for (let p = 2; p <= n; p++) {
        if (isPrime[p] === 1) {
            primes.push(p);
        }
    }

    return primes;
}

function main() {
    // Get the argument from process.argv[2]
    const arg = process.argv[2];

    if (arg === undefined) {
        // No argument provided, exit silently or handle error
        process.exit(0);
    }

    const n = parseInt(arg, 10);

    // Validate that the input is a valid number
    if (isNaN(n)) {
        process.exit(0);
    }

    const primes = getPrimesUpTo(n);

    // Print each prime on a new line
    if (primes.length > 0) {
        process.stdout.write(primes.join('\n') + '\n');
    }
}

// Run the script
main();