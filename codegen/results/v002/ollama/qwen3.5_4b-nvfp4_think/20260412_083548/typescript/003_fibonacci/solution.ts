import { argv } from 'process';

// Parse command-line argument - default to 0 if not provided
const N = parseInt(process.argv[2] || '0');

// Check if N < 1, exit without output
if (N < 1) {
  process.exit(0);
}

// Fibonacci sequence: starts with 1, 1, then each next is sum of previous two
let prev = 1;     // First number in sequence
let curr = 1;     // Second number in sequence

// Print first Fibonacci number (always 1 since N >= 1)
console.log(prev);

// Continue generating and printing numbers as long as curr <= N
while (curr <= N) {
  console.log(curr);
  
  // Calculate next Fibonacci number
  const next = prev + curr;
  
  // Update for next iteration
  prev = curr;
  curr = next;
}

// This stops when the next Fibonacci number would exceed N