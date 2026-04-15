import * as fs from 'fs';

// Get N from command line argument
const N = parseInt(process.argv[2], 10);

// If N is invalid or less than 2, there are no primes
if (isNaN(N) || N < 2) {
    process.exit(0);
}

// Sieve of Eratosthenes to find primes up to N
const isPrime = new Array(N + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let i = 2; i * i <= N; i++) {
    if (isPrime[i]) {
        for (let j = i * i; j <= N; j += i) {
            isPrime[j] = false;
        }
    }
}

// Print all primes, one per line
for (let i = 2; i <= N; i++) {
    if (isPrime[i]) {
        console.log(i);
    }
}