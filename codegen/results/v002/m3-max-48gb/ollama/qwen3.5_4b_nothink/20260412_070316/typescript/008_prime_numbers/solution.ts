import * as fs from 'fs';
import * as path from 'path';

// Get the argument N from command line
const argIndex = 2;
const N = parseInt(process.argv[argIndex], 10);

if (isNaN(N)) {
  process.exit(1);
}

if (N < 2) {
  process.exit(0);
}

// Sieve of Eratosthenes to find primes up to N
const MAX = N;
const isComposite = new Uint8Array(MAX + 1); // true means composite, false means prime
let count = 0;

// 0 and 1 are not prime, so initialize isComposite[0] and isComposite[1]
isComposite[0] = 1;
isComposite[1] = 1;

// Iterate from 2 to sqrt(MAX)
for (let p = 2; p * p <= MAX; p++) {
  if (!isComposite[p]) {
    // p is prime, mark its multiples as composite
    for (let multiple = p * p; multiple <= MAX; multiple += p) {
      isComposite[multiple] = 1;
    }
  }
}

// Collect primes
const primes: number[] = [];
for (let i = 2; i <= MAX; i++) {
  if (!isComposite[i]) {
    primes.push(i);
  }
}

// Print primes, one per line
for (const prime of primes) {
  console.log(prime);
}