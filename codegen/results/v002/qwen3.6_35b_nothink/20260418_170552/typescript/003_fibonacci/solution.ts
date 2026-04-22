const input = process.argv[2];
const N = parseInt(input, 10);

if (isNaN(N) || N < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

// Print the first two 1s
console.log(1);
if (N >= 1) {
  // We already printed the first 1, now check if we need another
  if (N >= b) {
    console.log(b);
  }
}

// Generate subsequent Fibonacci numbers
while (true) {
  const next = a + b;
  if (next > N) {
    break;
  }
  console.log(next);
  a = b;
  b = next;
}