/**
 * Fibonacci Generator
 * 
 * This script accepts a single integer argument N from the command line.
 * It prints all Fibonacci numbers up to and including the largest Fibonacci number
 * that does not exceed N.
 * 
 * Usage: tsx script.ts <N>
 */

function generateFibonacciUpToN(n: number): void {
  // If N is less than 1, print nothing
  if (n < 1) {
    return;
  }

  // Initialize the first two Fibonacci numbers
  let a = 1;
  let b = 1;

  // Print the first number if it's within range (1 <= N)
  if (a <= n) {
    console.log(a);
  }

  // Continue generating until the current number exceeds N
  while (b <= n) {
    console.log(b);
    
    // Calculate the next Fibonacci number
    const next = a + b;
    a = b;
    b = next;
  }
}

// Parse command line arguments
const args = process.argv.slice(2);

if (args.length === 0) {
  console.error("Error: Missing argument N.");
  process.exit(1);
}

const input = args[0];
const n = Number(input);

if (isNaN(n)) {
  console.error("Error: Argument N must be a valid integer.");
  process.exit(1);
}

generateFibonacciUpToN(n);