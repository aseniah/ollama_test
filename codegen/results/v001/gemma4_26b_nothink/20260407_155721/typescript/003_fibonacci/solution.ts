/**
 * This script calculates and prints Fibonacci numbers up to a given integer N.
 * Usage: tsx script.ts <N>
 */

function main(): void {
  // Read command line arguments
  const args = process.argv.slice(2);

  // Check if an argument is provided
  if (args.length === 0) {
    return;
  }

  // Parse the input argument as an integer
  const nStr = args[0];
  const n = parseInt(nStr, 10);

  // If N < 1, print nothing as per requirements
  if (isNaN(n) || n < 1) {
    return;
  }

  // Fibonacci sequence starts with 1, 1
  let a = 1;
  let b = 1;

  // Handle the first number in the sequence manually to facilitate the loop logic
  // (The loop will handle subsequent numbers)
  console.log(a);

  // If N is 1, we only print the first '1'. 
  // However, the sequence 1, 1, 2... implies we should check if the second 1 is <= N.
  if (b <= n) {
    console.log(b);
  } else {
    return;
  }

  // Calculate subsequent Fibonacci numbers
  while (true) {
    const next = a + b;

    // If the next number exceeds N, stop the sequence
    if (next > n) {
      break;
    }

    console.log(next);

    // Update pointers for next iteration
    a = b;
    b = next;
  }
}

// Execute the main function
main();