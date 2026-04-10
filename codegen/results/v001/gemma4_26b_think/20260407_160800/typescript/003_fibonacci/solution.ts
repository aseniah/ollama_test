/**
 * This script accepts a single integer argument N from the command line
 * and prints the Fibonacci sequence up to the largest Fibonacci number
 * that does not exceed N.
 * 
 * Sequence starts: 1, 1, 2, 3, 5, ...
 */

import process from 'process';

function main() {
  // Retrieve the first command-line argument
  const input = process.argv[2];

  // If no argument is provided, exit silently
  if (input === undefined) {
    return;
  }

  try {
    // Use BigInt to support arbitrarily large Fibonacci numbers 
    // without precision loss from standard 64-bit floats.
    const n = BigInt(input);

    // Requirement: If N < 1, print nothing.
    if (n < 1n) {
      return;
    }

    // Initialize the first two numbers of the Fibonacci sequence
    let current = 1n;
    let next = 1n;

    /**
     * We iterate as long as the current Fibonacci number is less than or equal to N.
     * We use a simple iterative approach to generate the sequence.
     * 
     * For N = 1:
     * Iteration 1: current = 1, next = 1. Print 1. next_val = 1 + 1 = 2. current becomes 1, next becomes 2.
     * Iteration 2: current = 1, next = 2. Print 1. next_val = 1 + 2 = 3. current becomes 2, next becomes 3.
     * Iteration 3: current = 2. 2 <= 1 is false. Loop terminates.
     * Result: 1, 1
     */
    while (current <= n) {
      // Print the current number followed by a newline
      process.stdout.write(current.toString() + '\n');

      // Calculate the next number in the sequence
      const temp = current + next;
      current = next;
      next = temp;
    }
  } catch (error) {
    // If the argument is not a valid integer, the BigInt constructor throws.
    // Based on requirements, we treat invalid input by printing nothing.
    return;
  }
}

// Execute the main function
main();