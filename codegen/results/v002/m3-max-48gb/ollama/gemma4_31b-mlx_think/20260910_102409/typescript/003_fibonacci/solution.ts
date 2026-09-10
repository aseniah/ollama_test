/**
 * The program accepts a single integer argument N from the command line.
 * It prints all Fibonacci numbers up to and including the largest
 * Fibonacci number that does not exceed N.
 * The sequence starts: 1, 1, 2, 3, 5, ...
 */

function solve() {
  // Read the argument from the command line
  const arg = process.argv[2];

  if (arg === undefined) {
    return;
  }

  try {
    // Use BigInt to support integers larger than 2^53 - 1
    const N = BigInt(arg);

    // If N < 1, print nothing as per requirements
    if (N < 1n) {
      return;
    }

    // Initialize Fibonacci sequence starting values
    // F(1) = 1, F(2) = 1
    let current = 1n;
    let next = 1n;

    // Loop and print numbers as long as the current Fibonacci number <= N
    while (current <= N) {
      process.stdout.write(current.toString() + '\n');
      
      // Calculate next term: F(n) = F(n-1) + F(n-2)
      const temp = current + next;
      current = next;
      next = temp;
    }
  } catch (e) {
    // If the argument is not a valid integer, we simply exit
    process.exit(0);
  }
}

solve();