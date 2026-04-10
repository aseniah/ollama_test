#!/usr/bin/env tsx
const args = process.argv.slice(2);

if (args.length !== 1) {
  console.error("Usage: npx esbuild --entry-file src/index.ts --target node");
  process.exit(1);
}

const N = parseInt(args[0], 10);

if (isNaN(N) || N < 2) {
  // No primes below 2
  process.exit(0);
}

// Simple Sieve of Eratosthenes
const primes = new Set<number>();

// 2 is the only even prime
primes.add(2);

// Check odd numbers from 3 to N
for (let candidate = 3; candidate <= N; candidate += 2) {
  let isPrime = true;
  const limit = Math.sqrt(candidate);
  
  // Only check primes up to sqrt(candidate)
  // Since we have all primes up to candidate in our set, we can check divisibility
  // But to optimize, we only need to check primes <= sqrt(candidate)
  // We iterate through our known primes
  for (const prime of primes) {
    if (prime * prime > candidate) {
      break;
    }
    if (candidate % prime === 0) {
      isPrime = false;
      break;
    }
  }
  
  if (isPrime) {
    primes.add(candidate);
  }
}

// Output primes in ascending order
for (const prime of primes) {
  console.log(prime);
}
process.exit(0);