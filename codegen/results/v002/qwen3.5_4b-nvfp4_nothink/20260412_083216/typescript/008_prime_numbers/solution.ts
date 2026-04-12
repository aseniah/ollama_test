import { argv } from "node:process";
import * as fs from "node:fs";

const args = argv.slice(1);
const n = parseInt(args[0], 10);

if (isNaN(n) || n < 2) {
  process.exit(0);
}

// Sieve of Eratosthenes to find all primes up to n
const isPrime = new Array<number>(n + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let p = 2; p <= Math.sqrt(n); p++) {
  if (isPrime[p]) {
    for (let multiple = p * p; multiple <= n; multiple += p) {
      isPrime[multiple] = false;
    }
  }
}

// Collect and print primes
const primes: number[] = [];
for (let i = 2; i <= n; i++) {
  if (isPrime[i]) {
    primes.push(i);
  }
}

if (primes.length > 0) {
  for (const prime of primes) {
    console.log(prime);
  }
} else {
  // Do not print anything if no primes exist
}