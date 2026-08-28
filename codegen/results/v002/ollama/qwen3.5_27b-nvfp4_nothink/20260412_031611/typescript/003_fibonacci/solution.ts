// Read command line argument (skip first two elements: node path and script name)
const input = process.argv[2];

if (!input) {
  // No argument provided, exit without printing anything
  process.exit(0);
}

const n = parseInt(input, 10);

// If N is not a valid number or is less than 1, print nothing
if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

while (a <= n) {
  console.log(a);
  const next = a + b;
  a = b;
  b = next;
}