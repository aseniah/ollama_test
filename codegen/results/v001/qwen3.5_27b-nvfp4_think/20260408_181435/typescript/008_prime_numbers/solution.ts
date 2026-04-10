#!/usr/bin/env tsx

const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
  process.exit(0);
}

function sieveOfEratosthenes(limit: number): number[] {
  const primes: number[] = [];
  const isPrime = new Array(limit + 1).fill(true);

  isPrime[0] = false;
  isPrime[1] = false;

  for (let i = 2; i <= limit; i++) {
    if (isPrime[i]) {
      primes.push(i);
      // Mark multiples starting from i*i, but only if within bounds
      const start = i * i;
      if (start <= limit) {
        for (let j = start; j <= limit; j += i) {
          isPrime[j] = false;
        }
      }
    }
  }

  return primes;
}

const primes = sieveOfEratosthenes(N);

for (const prime of primes) {
  console.log(prime);
}