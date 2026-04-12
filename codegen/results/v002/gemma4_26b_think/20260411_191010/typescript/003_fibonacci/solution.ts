/**
 * Problem:
 * Accept a single integer argument N from the command line.
 * Print all Fibonacci numbers up to and including the largest Fibonacci number
 * that does not exceed N. One number per line. Start the sequence: 1, 1, 2, 3, 5, ...
 * If N < 1, print nothing.
 */

import process from 'process';

/**
 * The solve function handles the logic of reading the input,
 * calculating the Fibonacci sequence, and printing the numbers.
 */
function solve(): void {
  // Retrieve the first command-line argument.
  const args = process.argv;
  const arg = args[2];

  // If no argument is provided, we exit without printing anything.
  if (arg === undefined) {
    return;
  }

  try {
    // We use BigInt to handle very large values of N without precision loss.
    // The problem specifies N is an integer.
    const n = BigInt(arg);

    // According to the requirements, if N < 1, we print nothing.
    if (n < 1n) {
      return;
    }

    /**
     * Initialize the Fibonacci sequence.
     * The sequence starts 1, 1, 2, 3, 5...
     * We use 'a' to represent the current Fibonacci number in the sequence 
     * and 'b' to represent the next number in the sequence.
     */
    let a = 1n;
    let b = 1n;

    // Since N >= 1, the first number (1) is always part of the sequence.
    console.log(a.toString());

    /**
     * We iterate through the sequence. 
     * In each step, 'b' represents the next Fibonacci number to be evaluated.
     * If 'b' does not exceed N, we print it and compute the subsequent number.
     */
    while (b <= n) {
      console.log(b.toString());
      
      // Calculate the next Fibonacci number: F(n) = F(n-1) + F(n-2)
      const next = a + b;
      
      // Advance the sequence pointers.
      a = b;
      b = next;
    }
  } catch (error) {
    /**
     * If the input argument cannot be parsed as a BigInt (e.g., non-integer input),
     * we stop execution.
     */
    process.exit(0);
  }
}

// Execute the solver.
solve();