/**
 * The program accepts a single integer N from the command line and prints
 * all Fibonacci numbers up to and including the largest Fibonacci number
 * that does not exceed N.
 * 
 * The Fibonacci sequence starts: 1, 1, 2, 3, 5, ...
 * If N < 1, nothing is printed.
 */

function main() {
  // Read N from command line arguments
  const input = process.argv[2];

  if (input === undefined) {
    return;
  }

  // Use BigInt to avoid precision issues with very large Fibonacci numbers
  let nValue: bigint;
  try {
    nValue = BigInt(input);
  } catch (e) {
    // If the input is not a valid integer string, we exit silently or could print an error.
    // Based on task requirements, we assume N is a valid integer.
    return;
  }

  // Requirement: If N < 1, print nothing.
  if (nValue < 1n) {
    return;
  }

  // Initialize the first two numbers of the sequence
  let current = 1n;
  let next = 1n;

  // The sequence starts 1, 1, 2, 3, 5...
  // We continue printing as long as the current Fibonacci number does not exceed N.
  while (current <= nValue) {
    // Print current number followed by a newline
    process.stdout.write(current.toString() + '\n');

    // Calculate the next term in the sequence
    const temp = current + next;
    current = next;
    next = temp;
  }
}

// Execute the main function
main();