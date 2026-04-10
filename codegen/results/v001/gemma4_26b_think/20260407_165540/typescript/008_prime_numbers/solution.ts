#!/usr/bin/env tsx

/**
 * This script implements the Sieve of Eratosthenes algorithm to find all 
 * prime numbers up to a given integer N provided via command line arguments.
 * 
 * Complexity:
 * - Time: O(N log log N)
 * - Space: O(N)
 */

function printPrimesUpTo(n: number): void {
  // If N is less than 2, there are no prime numbers.
  if (n < 2) {
    return;
  }

  // We use a Uint8Array to store the primality status.
  // Uint8Array is much more memory-efficient than a standard boolean array in Node.js.
  // index 'i' represents the number, value 1 means prime, 0 means not prime.
  const isPrime = new Uint8Array(n + 1).fill(1);
  
  // 0 and 1 are not prime numbers.
  isPrime[0] = 0;
  isPrime[1] = 0;

  const limit = Math.sqrt(n);

  for (let p = 2; p <= limit; p++) {
    if (isPrime[p] === 1) {
      // If p is prime, mark all multiples of p starting from p*p as not prime.
      // We start from p*p because smaller multiples would have been marked 
      // by smaller prime factors already.
      for (let multiple = p * p; multiple <= n; multiple += p) {
        isPrime[multiple] = 0;
      }
    }
  }

  // Buffering output to stdout to improve performance for large N.
  // Writing to stdout line-by-line (console.log) can be slow due to system call overhead.
  let buffer = "";
  const BUFFER_THRESHOLD = 16384; // 16KB buffer

  for (let i = 2; i <= n; i++) {
    if (isPrime[i] === 1) {
      buffer += i + "\n";

      // Periodically flush the buffer to stdout.
      if (buffer.length > BUFFER_THRESHOLD) {
        process.stdout.write(buffer);
        buffer = "";
      }
    }
  }

  // Final flush for any remaining content in the buffer.
  if (buffer.length > 0) {
    process.stdout.write(buffer);
  }
}

/**
 * Main execution block.
 * Reads the first command-line argument, parses it as an integer, and runs the sieve.
 */
function main(): void {
  const inputArg = process.argv[2];

  // If no argument is provided, exit quietly.
  if (inputArg === undefined) {
    return;
  }

  // Convert argument to integer.
  const n = parseInt(inputArg, 10);

  // If the argument is not a valid number, exit quietly.
  if (isNaN(n)) {
    return;
  }

  printPrimesUpTo(n);
}

// Execute the script.
main();