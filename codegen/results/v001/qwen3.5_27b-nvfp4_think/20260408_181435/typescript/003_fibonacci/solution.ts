const arg = process.argv[2];

if (!arg) {
  process.exit(0);
}

let n: bigint;

try {
  n = BigInt(arg);
} catch (_) {
  // If argument cannot be parsed as an integer, print nothing and exit.
  process.exit(0);
}

// If N < 1, print nothing.
if (n < 1n) {
  process.exit(0);
}

let a: bigint = 1n;
let b: bigint = 1n;

// Generate and print Fibonacci numbers up to n
while (a <= n) {
  console.log(a.toString());
  const temp = a + b;
  a = b;
  b = temp;
}