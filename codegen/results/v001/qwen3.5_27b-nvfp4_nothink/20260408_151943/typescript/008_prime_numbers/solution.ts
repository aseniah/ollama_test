import { readFileSync } from 'fs';
import * as path from 'path';

function main() {
  // Read command line arguments
  const args = process.argv.slice(2);

  if (args.length === 0) {
    return; // No argument provided, print nothing
  }

  const input = args[0];
  const n = parseInt(input, 10);

  if (isNaN(n)) {
    return; // Invalid integer, print nothing
  }

  if (n < 2) {
    return; // No primes less than 2
  }

  // Sieve of Eratosthenes
  // Create a boolean array "isPrime[0...n]" and initialize
  // all entries as true. A value in isPrime[i] will finally be false
  // if i is Not a prime, else true.
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

  // Print all prime numbers up to N
  for (let i = 2; i <= n; i++) {
    if (isPrime[i]) {
      console.log(i);
    }
  }
}

main();