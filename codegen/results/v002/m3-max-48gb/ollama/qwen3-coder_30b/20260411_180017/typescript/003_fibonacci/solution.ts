const N = parseInt(process.argv[2]);

if (isNaN(N) || N < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

// Print the first Fibonacci number if it's within the limit
if (a <= N) {
  console.log(a);
}

// Print the second Fibonacci number if it's within the limit
if (b <= N) {
  console.log(b);
}

// Generate and print subsequent Fibonacci numbers
while (true) {
  const next = a + b;
  if (next > N) {
    break;
  }
  console.log(next);
  a = b;
  b = next;
}