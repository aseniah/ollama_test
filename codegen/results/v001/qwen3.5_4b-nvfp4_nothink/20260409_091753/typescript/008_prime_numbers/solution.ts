import { argv } from "node:process";
import { stdout } from "node:util";

// Ensure we have the number as an argument
if (argv.length < 2) {
  console.error("Usage: tsx script.ts <N>");
  process.exit(1);
}

const n = parseInt(argv[2], 10);

if (isNaN(n) || n < 2) {
  // Nothing to print if N < 2
  process.exit(0);
}

// Helper function to check if a number is prime
function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2 || num === 3) return true;
  if (num % 2 === 0 || num % 3 === 0) return false;

  // Check divisors from 5 up to sqrt(num), skipping multiples of 2 and 3
  const limit = Math.sqrt(num);
  for (let i = 5; i <= limit; i += 6) {
    if (num % i === 0 || num % (i + 2) === 0) {
      return false;
    }
  }
  return true;
}

// Print primes up to N
let primeCount = 0; // To track number of output lines for potential future use if needed, though not required here.
const primes: number[] = [];

for (let i = 2; i <= n; i++) {
  if (isPrime(i)) {
    primes.push(i);
  }
}

// Output all found primes to stdout, one per line
for (const prime of primes) {
  stdout.write(`${prime}\n`);
}

process.exit(0);