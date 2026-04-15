/**
 * This program prints the Fibonacci sequence starting from 1, 1, 2...
 * up to and including the largest Fibonacci number that does not exceed N.
 */

function main() {
  // Read the first argument provided after the script execution command
  const arg = process.argv[2];

  if (arg === undefined) {
    return;
  }

  // Parse the input as an integer
  const N = parseInt(arg, 10);

  // If parsing failed (NaN) or N is less than 1, print nothing
  if (isNaN(N) || N < 1) {
    return;
  }

  let a = 1;
  let b = 1;

  // Print the first number of the sequence
  console.log(a);

  // Since the sequence starts 1, 1... we handle the second '1' 
  // and subsequent numbers in a loop.
  while (b <= N) {
    console.log(b);
    
    // Calculate the next number in the sequence
    const next = a + b;
    a = b;
    b = next;
  }
}

// Execute the program
main();