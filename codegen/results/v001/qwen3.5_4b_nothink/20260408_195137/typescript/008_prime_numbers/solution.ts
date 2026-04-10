const fs = require('fs');
const path = require('path');

// Parse command line argument
const argument = process.argv[2];
if (!argument || isNaN(parseInt(argument))) {
  process.exit(1);
}

const N = parseInt(argument);

// Handle negative numbers and numbers less than 2
// Note: By convention, primes are positive integers >= 2.
if (N < 2) {
  process.exit(0);
}

// Sieve of Eratosthenes algorithm to find primes up to N
const sieve = new Array(N + 1).fill(true);
sieve[0] = false;
sieve[1] = false;

const limit = Math.floor(Math.sqrt(N));

for (let i = 2; i <= limit; i++) {
  if (sieve[i]) {
    // Mark multiples of i as non-prime
    for (let j = i * i; j <= N; j += i) {
      sieve[j] = false;
    }
  }
}

// Collect and print prime numbers
const primes = [];
for (let i = 2; i <= N; i++) {
  if (sieve[i]) {
    primes.push(i);
  }
}

primes.forEach(prime => {
  console.log(prime);
});