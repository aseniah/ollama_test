// Accept a single integer argument N from the command line.
// Print all Fibonacci numbers up to and including the largest Fibonacci number
// that does not exceed N. One number per line. Start the sequence: 1, 1, 2, 3, 5, ...
// If N < 1, print nothing.

const args = process.argv.slice(2);
const nStr = args[0];

if (nStr === undefined) {
  // No argument provided, exit silently
  process.exit(0);
}

const n = parseInt(nStr, 10);

// If N < 1, print nothing
if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

// Print the first Fibonacci number (1)
if (a <= n) {
  console.log(a);
}

// Continue generating Fibonacci numbers
while (b <= n) {
  console.log(b);
  const next = a + b;
  a = b;
  b = next;
}