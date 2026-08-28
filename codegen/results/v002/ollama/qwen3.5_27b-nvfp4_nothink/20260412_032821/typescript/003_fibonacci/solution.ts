// Parse the command-line argument
const n = parseInt(process.argv[2], 10);

// If N is less than 1 or invalid, print nothing and exit
if (isNaN(n) || n < 1) {
  process.exit(0);
}

// Initialize the first two Fibonacci numbers in the sequence: 1, 1, ...
let a = 1;
let b = 1;

// Print the first number if it is <= N
if (a <= n) {
  console.log(a);
}

// Generate and print subsequent numbers while they are <= N
while (b <= n) {
  console.log(b);
  // Calculate the next Fibonacci number
  const next = a + b;
  a = b;
  b = next;
}