import * as fs from 'fs';
import * as path from 'path';

// Read command-line argument
const arg = process.argv[2];
if (!arg) {
  // No argument provided
  process.exit(0);
}

// Parse N as integer
let n: number;
try {
  n = parseInt(arg, 10);
} catch {
  // Invalid argument
  process.exit(0);
}

// If N < 1, print nothing
if (n < 1) {
  process.exit(0);
}

// Generate Fibonacci numbers up to N
let a = 1;
let b = 1;

// If n >= 1, print the first 1
if (n >= 1) {
  console.log(a);
  
  // Continue generating Fibonacci numbers
  while (true) {
    const next = a + b;
    if (next > n) {
      break;
    }
    a = b;
    b = next;
    console.log(a);
  }
}