#!/usr/bin/env tsx

function main() {
  // Get the command-line argument (process.argv[0] is the runtime, process.argv[1] is the script name)
  const input = process.argv[2];

  if (!input) {
    // No argument provided
    return;
  }

  const n = parseInt(input, 10);

  if (isNaN(n) || n < 1) {
    // Invalid input or less than 1, print nothing
    return;
  }

  // Fibonacci sequence: 1, 1, 2, 3, 5, ...
  let a = 1;
  let b = 1;

  while (a <= n) {
    console.log(a);
    const next = a + b;
    a = b;
    b = next;
  }
}

main();