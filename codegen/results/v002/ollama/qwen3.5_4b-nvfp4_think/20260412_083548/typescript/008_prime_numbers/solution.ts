import { readFileSync } from 'fs';

const N = parseInt(process.argv[2] || "0");

function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return num === 2;

  for (let i = 3; i * i <= num; i += 2) {
    if (num % i === 0) return false;
  }

  return true;
}

const primes: number[] = [];
for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    primes.push(i);
  }
}

primes.forEach(prime => console.log(prime.toString()));