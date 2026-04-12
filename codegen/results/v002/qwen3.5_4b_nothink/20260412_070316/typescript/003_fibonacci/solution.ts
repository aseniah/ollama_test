import { readFileSync } from 'fs';

// Read N from command-line argument
const N = parseInt(process.argv[2], 10);

// If N < 1, print nothing and exit
if (N < 1) {
  process.exit(0);
}

// Generate Fibonacci numbers until the next one exceeds N
let a = 1;
let b = 1;

do {
  console.log(a);
  const sum = a + b;
  if (sum > N) {
    break;
  }
  a = b;
  b = sum;
} while (true);