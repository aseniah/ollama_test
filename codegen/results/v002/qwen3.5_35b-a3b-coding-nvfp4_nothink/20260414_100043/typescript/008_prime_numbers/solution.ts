import * as fs from 'fs';

// Read N from command line arguments
const arg = process.argv[2];
if (!arg) {
  process.exit(0);
}

const N = parseInt(arg, 10);

if (isNaN(N) || N < 2) {
  // No primes up to N
  process.exit(0);
}

// Sieve of Eratosthenes to find primes up to N
const isPrime = new Array<boolean>(N + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let i = 2; i * i <= N; i++) {
  if (isPrime[i]) {
    for (let j = i * i; j <= N; j += i) {
      isPrime[j] = false;
    }
  }
}

// Print all primes up to N
for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    console.log(i);
  }
}