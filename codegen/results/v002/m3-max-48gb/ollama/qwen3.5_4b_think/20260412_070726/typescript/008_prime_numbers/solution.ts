#!/usr/bin/env tsx
import { parseInt } from 'util';
// No external modules needed - using only Node.js built-ins

function isPrime(n: number): boolean {
  if (n < 2) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;
  for (let i = 3; i <= Math.sqrt(n); i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

function main(): void {
  const arg = process.argv[2];
  if (!arg || isNaN(parseInt(arg))) {
    process.stderr.write('Error: Invalid input. Please provide an integer N.\n');
    process.exit(1);
  }

  const N = parseInt(arg, 10);
  if (isNaN(N) || N < 0) {
    process.stderr.write('Error: N must be a non-negative integer.\n');
    process.exit(1);
  }

  const primes: number[] = [];
  for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
      primes.push(i);
    }
  }

  for (const prime of primes) {
    console.log(prime);
  }
}

main();