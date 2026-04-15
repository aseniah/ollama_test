// Fibonacci sequence generator up to N
// Usage: tsx script.ts <N>
// Where N is a single integer argument.

const args = process.argv.slice(2);

if (args.length !== 1) {
  // No argument provided, do nothing
  process.exit(0);
}

const nStr = args[0];
const n = parseInt(nStr, 10);

// Validate input
if (isNaN(n) || !Number.isInteger(n)) {
  process.exit(0);
}

if (n < 1) {
  // If N < 1, print nothing
  process.exit(0);
}

// Generate Fibonacci numbers starting with 1, 1, 2, 3, 5, ...
let a = 1;
let b = 1;

while (a <= n) {
  console.log(a);
  const next = a + b;
  a = b;
  b = next;
}