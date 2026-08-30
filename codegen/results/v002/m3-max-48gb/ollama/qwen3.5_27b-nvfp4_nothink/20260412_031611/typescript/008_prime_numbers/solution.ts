/**
 * Prints all prime numbers up to and including N.
 * 
 * Usage: tsx script.ts <N>
 */

function sieveOfEratosthenes(limit: number): boolean[] {
  if (limit < 2) return [];

  // Initialize a boolean array where index represents the number
  // true means the number is potentially prime, false means composite
  const isPrime = new Array(limit + 1).fill(true);
  isPrime[0] = false;
  isPrime[1] = false;

  // Iterate from 2 up to sqrt(limit)
  for (let i = 2; i * i <= limit; i++) {
    if (isPrime[i]) {
      // Mark multiples of i starting from i*i as non-prime
      for (let j = i * i; j <= limit; j += i) {
        isPrime[j] = false;
      }
    }
  }

  return isPrime;
}

function main() {
  // Check if an argument was provided
  if (process.argv.length < 3) {
    // If no argument, print nothing as per requirements (no primes to check)
    return;
  }

  const input = process.argv[2];
  const n = parseInt(input, 10);

  // Validate input
  if (isNaN(n)) {
    return;
  }

  // If N is less than 2, there are no primes
  if (n < 2) {
    return;
  }

  const primesFlags = sieveOfEratosthenes(n);

  // Iterate and print indices that are marked as prime
  for (let i = 2; i <= n; i++) {
    if (primesFlags[i]) {
      console.log(i);
    }
  }
}

main();