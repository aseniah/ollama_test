import { readFileSync } from 'fs';

// Read command line arguments
// process.argv[0] is the node/ts executable
// process.argv[1] is the script name
// process.argv[2] is the first argument provided by the user
const inputArg = process.argv[2];

if (!inputArg) {
  // No argument provided, exit silently or handle as per requirements (print nothing)
  process.exit(0);
}

const N = parseInt(inputArg, 10);

if (isNaN(N) || N < 2) {
  // If N is not a valid number or less than 2, there are no primes.
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

// Collect and print primes
const primes: number[] = [];
for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    primes.push(i);
  }
}

if (primes.length > 0) {
  console.log(primes.join('\n'));
}