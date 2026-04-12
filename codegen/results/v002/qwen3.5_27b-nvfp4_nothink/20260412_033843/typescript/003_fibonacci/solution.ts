#!/usr/bin/env tsx

/**
 * This TypeScript program accepts a single integer argument N from the command line.
 * It prints all Fibonacci numbers up to and including the largest Fibonacci number
 * that does not exceed N. The sequence starts with 1, 1, 2, 3, 5, ...
 * If N < 1, it prints nothing.
 */

function printFibonacciUpTo(limit: number): void {
  if (limit < 1) {
    return;
  }

  let a = 1;
  let b = 1;

  // Print the first Fibonacci number
  console.log(a);

  // If limit is 1, we stop here because the next number is also 1, 
  // but the sequence definition 1, 1, 2... implies we should list both if N >= 1?
  // The prompt says "Start the sequence: 1, 1, 2, 3, 5, ..."
  // If N=1, the numbers <= 1 are 1 and 1. So we print both.
  
  // Loop to generate subsequent numbers
  while (b <= limit) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
  }
}

// Main execution logic
function main(): void {
  // Check if an argument is provided
  if (process.argv.length < 3) {
    // No argument provided, print nothing as per "If N < 1" logic or simply do nothing.
    // The problem says "Accept a single integer argument", implying it should be there.
    // If not provided, we treat it as an invalid case similar to N < 1 behavior (print nothing).
    return;
  }

  const input = process.argv[2];
  const n = Number(input);

  // Check if input is a valid integer
  if (!Number.isInteger(n)) {
    return;
  }

  printFibonacciUpTo(n);
}

main();