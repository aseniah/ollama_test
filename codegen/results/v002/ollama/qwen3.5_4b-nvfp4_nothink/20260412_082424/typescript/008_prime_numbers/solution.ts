import { parseArgs } from 'node:util';

const { name } = parseArgs({
  args: process.argv.slice(2),
  option: { n: { type: 'number', default: undefined } },
});

if (!name) {
  console.error('Usage: node <script.ts> <N>');
  process.exit(1);
}

const N = parseInt(name, 10);

if (isNaN(N) || N < 2) {
  // If N is less than 2, there are no prime numbers.
  // We print nothing as per the requirement.
  return;
}

// Sieve of Eratosthenes to find primes up to N
const limit = Math.floor(N);
const sieve = new Array<boolean>(limit + 1).fill(true);
sieve[0] = false;
sieve[1] = false;

for (let i = 2; i * i <= limit; i++) {
  if (sieve[i]) {
    for (let j = i * i; j <= limit; j += i) {
      sieve[j] = false;
    }
  }
}

const primes: number[] = [];
for (let i = 2; i <= limit; i++) {
  if (sieve[i]) {
    primes.push(i);
  }
}

// Print all primes found, one per line
primes.forEach(p => console.log(p));