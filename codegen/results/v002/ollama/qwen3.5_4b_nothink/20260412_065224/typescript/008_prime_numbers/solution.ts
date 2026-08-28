import * as fs from 'fs';

const [ , Nstr ] = process.argv;
const N = parseInt(Nstr, 10);

if (isNaN(N) || N < 2) {
  // No primes less than 2, so print nothing
  process.exit(0);
}

const primes: number[] = [];

// Sieve of Eratosthenes approach
// We can build a boolean array marking non-primes
const isPrime: boolean[] = new Array(N + 1).fill(true);

for (let i = 2; i * i <= N; i++) {
  if (isPrime[i]) {
    for (let j = i * i; j <= N; j += i) {
      isPrime[j] = false;
    }
  }
}

for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    primes.push(i);
  }
}

// Print all primes, one per line
for (const p of primes) {
  console.log(p);
}