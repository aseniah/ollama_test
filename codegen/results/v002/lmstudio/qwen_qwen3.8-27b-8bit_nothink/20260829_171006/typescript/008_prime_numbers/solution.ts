#!/usr/bin/env tsx

const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
    // No primes below 2, print nothing
    process.exit(0);
}

// Sieve of Eratosthenes
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

const output: string[] = [];
for (let i = 2; i <= N; i++) {
    if (isPrime[i]) {
        output.push(i.toString());
    }
}

if (output.length > 0) {
    console.log(output.join('\n'));
}