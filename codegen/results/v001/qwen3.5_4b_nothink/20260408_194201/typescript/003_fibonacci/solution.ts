import { argv } from 'process';

// Read N from command line argument
const N = parseInt(argv[2], 10);

// Check if N is a valid positive integer
if (Number.isNaN(N) || N < 1) {
  // If N < 1 or invalid, print nothing
  process.exit(0);
}

// Generate Fibonacci numbers up to N
let a = 1;
let b = 1;

// Print the first two numbers (1, 1) if they are within range
if (a <= N) {
  console.log(a);
}
if (b <= N) {
  console.log(b);
}

// Generate subsequent Fibonacci numbers
while (true) {
  const next = a + b;
  
  // If next is greater than N, stop
  if (next > N) {
    break;
  }
  
  console.log(next);
  
  // Update a and b for the next iteration
  a = b;
  b = next;
}