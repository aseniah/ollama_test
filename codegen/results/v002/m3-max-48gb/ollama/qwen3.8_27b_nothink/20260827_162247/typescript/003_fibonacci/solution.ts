const N = parseInt(process.argv[2] || "", 10);

if (isNaN(N) || N < 1) {
  // Print nothing
  process.exit(0);
}

const results: number[] = [];
let a = 1;
let b = 1;

results.push(a);
if (N >= a) {
  // We've already printed the first 1
}

// Start with the sequence: 1, 1, 2, 3, 5, ...
// We need to generate Fibonacci numbers up to and including the largest one that does not exceed N.
// The sequence starts with 1, 1, 2, 3, 5, 8, ...

let fib1 = 1;
let fib2 = 1;

// Print the first two ones
if (fib1 <= N) {
  console.log(fib1);
}
if (fib2 <= N) {
  console.log(fib2);
}

// Now generate subsequent Fibonacci numbers
let next = fib1 + fib2;
while (next <= N) {
  console.log(next);
  fib1 = fib2;
  fib2 = next;
  next = fib1 + fib2;
}