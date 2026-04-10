/**
 * This script takes a single integer N from the command line
 * and prints all prime numbers up to and including N, one per line.
 * 
 * Usage: tsx script.ts <N>
 */

import process from 'node:process';

/**
 * Sieve of Eratosthenes algorithm to find all primes up to N.
 * Complexity: O(N log log N)
 * 
 * @param n The upper bound (inclusive)
 * @returns An array of prime numbers up to n
 */
function getPrimesUpTo(n: number): number[] {
  if (n < 2) return [];

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
      // Multiples less than p*p are already marked
      for (let i = p * p; i <= n; i += p) {
        isPrime[i] = false;
      }
    }
  }

  const primes: number[] = [];
  for (let p = 2; p <= n; p++) {
    if (isPrime[p]) {
      primes.push(p);
    }
  }
  return primes;
}

function main() {
  const args = process.argv.slice(2);

  if (args.length === 0) {
    // No argument provided, exit silently or handle error
    return;
  }

  const input = args[0];
  const n = parseInt(input, 10);

  // Validate input
  if (isNaN(n)) {
    return;
  }

  // Find primes
  const primes = getPrimesUpTo(n);

  // Print each prime on a new line
  if (primes.length > 0) {
    process.stdout.write(primes.join('\n') + '\n');
  }
}

// Execute the script
main();