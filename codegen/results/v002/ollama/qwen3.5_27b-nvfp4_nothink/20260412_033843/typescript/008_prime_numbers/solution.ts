#!/usr/bin/env tsx

/**
 * Script: Find all prime numbers up to N (inclusive) using the Sieve of Eratosthenes.
 * Input: A single integer argument from command line (process.argv[2]).
 * Output: Prints each prime number on a new line to stdout.
 */

function main(): void {
  const input = process.argv[2];

  if (input === undefined) {
    return;
  }

  const n = parseInt(input, 10);

  // If N is less than 2, there are no primes, so print nothing and exit.
  if (isNaN(n) || n < 2) {
    return;
  }

  // Sieve of Eratosthenes to find all primes up to n
  // Create a boolean array "isPrime" where index represents the number.
  // Initialize all entries as true.
  const isPrime = new Array<boolean>(n + 1).fill(true);

  // 0 and 1 are not prime numbers
  isPrime[0] = false;
  isPrime[1] = false;

  // Only need to check up to sqrt(n)
  for (let i = 2; i * i <= n; i++) {
    if (isPrime[i]) {
      // If i is prime, then mark all multiples of i as non-prime
      // Start from i*i because smaller multiples would have been handled already
      for (let j = i * i; j <= n; j += i) {
        isPrime[j] = false;
      }
    }
  }

  // Collect and print all prime numbers
  const primes: string[] = [];
  for (let i = 2; i <= n; i++) {
    if (isPrime[i]) {
      primes.push(i.toString());
    }
  }

  // Print all found primes, one per line
  if (primes.length > 0) {
    console.log(primes.join('\n'));
  }
}

main();