// Solve the "Prime Numbers" task
import { stdout } from 'node:process';

function main() {
  const args = process.argv.slice(1);

  if (args.length < 1) {
    console.error('Usage: node solution.ts <N>');
    process.exit(1);
  }

  const N = parseInt(args[0], 10);
  if (Number.isNaN(N) || N < 0) {
    console.error('Error: Invalid integer argument.');
    process.exit(1);
  }

  // A simple sieve of Eratosthenes up to N.
  const isPrime = new Array<number | undefined>(N + 1).fill(undefined);

  // Mark all numbers as potentially prime initially (default true, but we track via negation later if needed).
  // Actually, simpler: just mark non-primes.
  for (let i = 2; i <= N; i++) {
    isPrime[i] = 0; // assume not prime until proven otherwise
  }

  for (let p = 2; p * p <= N; p++) {
    if (!isPrime[p]) {
      for (let multiple = p * p; multiple <= N; multiple += p) {
        isPrime[multiple] = 0; // mark as composite
      }
    }
  }

  const primes: number[] = [];

  for (let num = 2; num <= N; num++) {
    if (isPrime[num] === undefined || isPrime[num] !== 0) {
      primes.push(num);
    }
  }

  // If there are primes, print them. Otherwise, print nothing.
  for (const prime of primes) {
    stdout.write(`${prime}\n`);
  }
}

main();