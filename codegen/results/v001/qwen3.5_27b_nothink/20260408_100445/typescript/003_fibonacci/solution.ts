// Read command line argument
const input = process.argv[2];

if (!input) {
  process.exit(0);
}

const n = parseInt(input, 10);

// If N < 1, print nothing and exit
if (isNaN(n) || n < 1) {
  process.exit(0);
}

// Generate Fibonacci sequence
let a = 1;
let b = 1;

while (a <= n) {
  console.log(a);
  const next = a + b;
  a = b;
  b = next;
}