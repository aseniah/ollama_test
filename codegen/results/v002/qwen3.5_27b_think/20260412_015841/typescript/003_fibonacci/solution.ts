// Accept a single integer argument N from the command line.
// process.argv[0] is the node executable, process.argv[1] is the script name.
// process.argv[2] is the first argument.
const input = process.argv[2];

// Check if input exists and is a valid number
if (!input) {
  process.exit(0);
}

const n = parseInt(input, 10);

// If N < 1 (or is NaN), print nothing.
if (isNaN(n) || n < 1) {
  process.exit(0);
}

// Initialize Fibonacci sequence
// Start with 1, 1
let a: number = 1;
let b: number = 1;

// Print all Fibonacci numbers up to and including N
while (a <= n) {
  console.log(a);
  
  const next = a + b;
  a = b;
  b = next;
}