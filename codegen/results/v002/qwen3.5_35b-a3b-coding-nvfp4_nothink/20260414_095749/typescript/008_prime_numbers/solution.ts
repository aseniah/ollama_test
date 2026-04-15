/**
 * Accepts a single integer argument N from the command line.
 * Prints all prime numbers up to and including N, one per line.
 * Uses the Sieve of Eratosthenes for efficient prime generation.
 */

const args = process.argv.slice(2);

if (args.length !== 1) {
  process.exit(0); // No input provided, print nothing
}

const nStr = args[0];
const n = parseInt(nStr, 10);

// Validate input
if (isNaN(n) || n < 2) {
  // No primes up to N (if N < 2), print nothing
  process.exit(0);
}

// Sieve of Eratosthenes
const isPrime = new Array<boolean>(n + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

const limit = Math.floor(Math.sqrt(n));

for (let i = 2; i <= limit; i++) {
  if (isPrime[i]) {
    for (let j = i * i; j <= n; j += i) {
      isPrime[j] = false;
    }
  }
}

// Collect and print primes
const output: string[] = [];
for (let i = 2; i <= n; i++) {
  if (isPrime[i]) {
    output.push(i.toString());
  }
}

if (output.length > 0) {
  process.stdout.write(output.join('\n') + '\n');
}