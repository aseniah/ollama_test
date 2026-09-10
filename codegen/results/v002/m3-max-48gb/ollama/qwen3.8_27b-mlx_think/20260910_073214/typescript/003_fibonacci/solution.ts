const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
  // N < 1, print nothing
  process.exit(0);
}

const results: string[] = [];
let a = 1;
let b = 1;

// Print all Fibonacci numbers <= N
// Sequence: 1, 1, 2, 3, 5, 8, ...
let first = true;
let second = true;

// First 1
if (1 <= N) {
  results.push("1");
  first = false;
}

// Second 1
if (1 <= N) {
  results.push("1");
  second = false;
}

// Generate subsequent Fibonacci numbers
let prev = 1;
let curr = 1;
while (true) {
  const next = prev + curr;
  if (next > N) break;
  results.push(next.toString());
  prev = curr;
  curr = next;
}

console.log(results.join("\n"));