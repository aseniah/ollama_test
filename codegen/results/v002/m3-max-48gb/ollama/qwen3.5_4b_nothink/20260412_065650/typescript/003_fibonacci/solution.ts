// Read N from command line arguments
const N = parseInt(process.argv[2], 10);

// If N is less than 1, print nothing and exit
if (N < 1) {
  process.exit(0);
}

// Generate Fibonacci numbers until the next one would exceed N
let a = 1;
let b = 1;

// Print the first Fibonacci number
if (a <= N) {
  console.log(a);
}

// Generate and print subsequent Fibonacci numbers
while (true) {
  const next = a + b;
  if (next > N) {
    break;
  }
  b = a;
  a = next;
  if (a <= N) {
    console.log(a);
  }
}