/**
 * This script accepts a single integer argument N from the command line.
 * It prints the Fibonacci sequence starting with 1, 1, 2, 3, 5, ...
 * up to the largest Fibonacci number that does not exceed N.
 * If N < 1, nothing is printed.
 */

function runFibonacci() {
  // Extract the first command-line argument
  const arg = process.argv[2];

  // If no argument is provided, terminate
  if (arg === undefined) {
    return;
  }

  try {
    // Convert argument to BigInt to support extremely large values of N
    const N = BigInt(arg);

    // Requirement: If N < 1, print nothing.
    if (N < 1n) {
      return;
    }

    /**
     * We initialize the sequence with the first two terms: 1 and 1.
     * To ensure we print the duplicate '1' in the sequence (1, 1, 2, ...),
     * we use a state where 'curr' is the current term to be printed
     * and 'next' is the following term in the sequence.
     */
    let curr = 1n;
    let next = 1n;

    // Continue the sequence as long as the current term does not exceed N
    while (curr <= N) {
      // Print the current Fibonacci number followed by a newline
      process.stdout.write(curr.toString() + '\n');

      /**
       * Calculate the next term in the sequence.
       * The next term in the Fibonacci sequence is the sum of the current
       * and the subsequent term.
       * 
       * For the sequence 1, 1, 2, 3, 5...
       * Iteration 1: curr=1, next=1. Print 1. New curr=1, new next=2.
       * Iteration 2: curr=1, next=2. Print 1. New curr=2, new next=3.
       * Iteration 3: curr=2, next=3. Print 2. New curr=3, new next=5.
       */
      const temp = curr + next;
      curr = next;
      next = temp;
    }
  } catch (error) {
    // If the input cannot be parsed as a BigInt, we catch the error and exit silently.
    // This handles non-integer strings or malformed input.
  }
}

// Execute the main function
runFibonacci();